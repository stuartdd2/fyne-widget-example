package main

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	PI_180_F64           = float64(math.Pi / 180.0) // Used to convert degrees to radians
	DEG_PER_TICK_INT     = 6                        // 60 clock hand positions on a 360 degree face. 6 degrees per pos
	HOUR_TO_60_INT       = 5                        // 5 clock positions per hour for the hour hand
	MIN_IN_5DEG_F64      = float64(5.0 / 60.0)      // positions between numerals for minutes of the hour hand
	SEC_HAND_LEN_F32     = 0.77                     // adjusted length of the second hand
	MIN_HAND_LEN_F32     = 0.65                     // adjusted length of the minute hand
	HOUR_HAND_LEN_F32    = 0.55                     // adjusted length of the hour hand
	CENTER_SIZE_FACT_F32 = 15.0                     // Radius of center circle scaled from radius of clock face
)

//
// These do not add to the Widget they are here to
// indicate conformance to specific interfaces.
// If you break the widget implementation of an interface
// the respective line below will indicate an error.
// This is really usefull for widget development
//
var _ fyne.WidgetRenderer = (*widgetClockRenderer)(nil)
var _ fyne.CanvasObject = (*WidgetClock)(nil)
var _ fyne.Widget = (*WidgetClock)(nil)
var _ fyne.Tappable = (*WidgetClock)(nil)
var _ fyne.SecondaryTappable = (*WidgetClock)(nil)

//
// The Renderer (fyne.WidgetRenderer interface) is responsible for ALL
// drawing, scaling, moving, animating etc
//
type widgetClockRenderer struct {
	widget     *WidgetClock      // Reference to the widget so we can access state information
	size       fyne.Size         // The current size of the widget. Set when Layout is called
	background *canvas.Rectangle // A rectangle that forms the background
	centCirc   *canvas.Circle    // A circle at the center of the clock dispay
	sHand      *canvas.Line      // The second hand
	mHand      *canvas.Line      // The minute hand
	hHand      *canvas.Line      // The hour hand
	h, m, s    int               // The hour, minute and second currently shown
}

func newWidgetClockRenderer(w *WidgetClock) *widgetClockRenderer {
	r := &widgetClockRenderer{
		widget: w,
		size:   w.MinSize(),
		background: &canvas.Rectangle{
			FillColor: w.BackgroundColor,
		},
	}
	r.widget.Image.FillMode = canvas.ImageFillContain
	r.mHand = canvas.NewLine(r.widget.MHandColor)
	r.mHand.StrokeWidth = 3
	r.hHand = canvas.NewLine(r.widget.HHandColor)
	r.hHand.StrokeWidth = 3
	r.sHand = canvas.NewLine(r.widget.SHandColor)
	r.sHand.StrokeWidth = 1
	r.centCirc = canvas.NewCircle(r.widget.CenterColor)
	// Update the time at the start
	r.updateTime()
	//
	// Start the bacground animation. One process per second
	//
	tick := time.NewTicker(time.Second)
	go func() {
		// Run forever in the background
		for {
			// If running updatethe time
			if r.widget.running {
				r.updateTime()
			}
			// Lay everything out
			r.Layout(r.size)
			// Redraw the widget
			canvas.Refresh(r.widget)
			// Wait for the next tick
			<-tick.C
		}
	}()
	// Return the renderer
	return r
}

//
// From the WidgetRenderer interface.
// For the given size. Layout (rearrange) all of the visible components.
// This also updates the position of the clock hands.
//
func (r *widgetClockRenderer) Layout(s fyne.Size) {
	// Somtimes Layout is called with invalid or 0 width or height.
	// If invalid then do nothing
	if s.Width <= 1 || s.Height <= 1 {
		return
	}

	r.size = s // Save the size
	// Calc image size and position.
	shrink := r.widget.ImageShrink
	offset := shrink / 2
	imageSize := fyne.Size{Width: s.Width - shrink, Height: s.Height - shrink}
	// Set image size and position.
	r.widget.Image.Resize(imageSize)
	r.widget.Image.Move(fyne.Position{X: offset, Y: offset})

	// Set background size and position.
	r.background.Resize(s)
	r.background.FillColor = r.widget.BackgroundColor

	// Calc center circle size and position.
	imageCent := fyne.Position{X: (imageSize.Width / 2) + offset, Y: (imageSize.Height / 2) + offset}
	centCircW := imageSize.Width / CENTER_SIZE_FACT_F32
	centCircH := imageSize.Height / CENTER_SIZE_FACT_F32
	// Set center circle size and position.
	r.centCirc.Resize(fyne.Size{Width: centCircW, Height: centCircH})
	r.centCirc.Move(fyne.Position{X: imageCent.X - (centCircW / 2), Y: imageCent.Y - (centCircH / 2)})

	// Set colours of all components
	r.centCirc.FillColor = r.widget.CenterColor
	r.sHand.StrokeColor = r.widget.SHandColor
	r.mHand.StrokeColor = r.widget.MHandColor
	r.hHand.StrokeColor = r.widget.HHandColor

	// Calc the full length of the line used to the clock hands
	// Uses the lowest width or height of tthe image (radius of the clock face)
	var lineLen float32
	if imageSize.Width > imageSize.Height {
		lineLen = imageSize.Height / 2
	} else {
		lineLen = imageSize.Width / 2
	}
	// Update each line (clock hand)
	// HAND_LEN_F32 is used to scale each hand to different lengths
	r.updateClockHand(imageCent, lineLen*SEC_HAND_LEN_F32, r.s, r.sHand)
	r.updateClockHand(imageCent, lineLen*MIN_HAND_LEN_F32, r.m, r.mHand)
	r.updateClockHand(imageCent, lineLen*HOUR_HAND_LEN_F32, hourToClock60(r.h, r.m), r.hHand)
}

//
// From the WidgetRenderer interface.
// Return the minimum size of the widget. Get this from the widget!
//
func (r *widgetClockRenderer) MinSize() fyne.Size {
	return r.widget.MinSize()
}

//
// From the WidgetRenderer interface.
// Does not seem to be called! It is required for interface complience
//
func (r *widgetClockRenderer) Refresh() {}

//
// From the WidgetRenderer interface.
// Return a list of CanvasObjects that will require display (rendering)
// The order is critical. The last object in the list is the lat to be drawn
//
func (r *widgetClockRenderer) Objects() []fyne.CanvasObject {
	o := make([]fyne.CanvasObject, 0)
	o = append(o, r.background, r.widget.Image, r.sHand, r.hHand, r.mHand, r.centCirc)
	return o
}

//
// From the WidgetRenderer interface.
// Called when the rendered is destroyed and the memory released.
//  This is where you clean up your mess!
//  Nothing to do at the moment!
//
func (r *widgetClockRenderer) Destroy() {
}

//
// Update the time using the system clock
//  Do not call this if the clock is not running
//
func (r *widgetClockRenderer) updateTime() {
	r.h = time.Now().Hour()
	r.m = time.Now().Minute()
	r.s = time.Now().Second()
}

//
// Detirmine the position for the hour hand.
//    The are 60 positions corresponding to 360 degrees on the clock face.
//    The hour hand takes 12 hours plus additional degrees for additional minutes.
//    This means the hour hand moves between the clock numerals
//
func hourToClock60(hr, min int) int {
	return ((hr % 12) * HOUR_TO_60_INT) + int(float64(min)*MIN_IN_5DEG_F64)
}

//
// The are 60 positions corresponding to 360 degrees on the clock face.
// Given the position and length we use trigonometry to derive the end of the line
//    x = len * cos(theta)
//    y = len * sin(theta)
// The start of the line is the center of the clock face .
//   Note the math api uses float64 but every thing else is in float32 so
//   some conversion must take place.
//   Radians (float64) are used by math Cos and Sine functions.
//
func (r *widgetClockRenderer) updateClockHand(cent fyne.Position, lineLen float32, position int, line *canvas.Line) {
	radians := (float64((position*DEG_PER_TICK_INT)-90) * PI_180_F64)
	xx := cent.X + lineLen*float32(math.Cos(radians))
	yy := cent.Y + lineLen*float32(math.Sin(radians))
	line.Position1 = cent
	line.Position2 = fyne.Position{X: xx, Y: yy}
}

//
// The state of the widget.
// The properties here can be updated dynamically to change the
// apperance of the clock.
//
type WidgetClock struct {
	widget.BaseWidget               // Inherit from BaseWidget
	running           bool          // Clock time is updating (hands moving)
	minSize           fyne.Size     // The minimum size of the widget
	Image             *canvas.Image // An image displayed above the background (clock face)
	ImageShrink       float32       // Make the image smaller by this amount (provides a boarder)
	HHandColor        color.Color   // The hour hand colour
	MHandColor        color.Color   // The minute hand colour
	SHandColor        color.Color   // The second hand colour
	CenterColor       color.Color   // The blob at the clock center colour
	BackgroundColor   color.Color   // The background colour
}

//
// Create a Clock Widget and Extend the BaseWidget
// Minimum parameters provided.
//   The amount to shrink the image (provides a boarder)
//	 The image. This should be the clock face
//   W and H used to set the minimum Width and Height of the widget
//
func NewWidgetClock(imageShrink float32, image *canvas.Image, W, H float32) *WidgetClock {
	w := &WidgetClock{
		minSize:     fyne.Size{Width: W, Height: H},
		ImageShrink: imageShrink,
		Image:       image,
		// Define the default values other properties
		running:         false,
		MHandColor:      theme.PrimaryColorNamed("green"),
		SHandColor:      theme.PrimaryColorNamed("red"),
		HHandColor:      theme.PrimaryColorNamed("blue"),
		CenterColor:     theme.PrimaryColorNamed("green"),
		BackgroundColor: color.White,
	}
	w.ExtendBaseWidget(w)
	return w
}

//
// Part of the widget interface. Provided a reference to the Renderer
// Do not hang on to the reference to the Renderer. fyne may stop using
// it and call again for a new one. fyne also caches the reference.
// The Renderer gets the widget state from a passed in reference to the widget.
//
func (w *WidgetClock) CreateRenderer() fyne.WidgetRenderer {
	return newWidgetClockRenderer(w)
}

//
// Start (true) or Stop (false) the clock hands.
// The cock hands are always updated (in case the widget is resized)
// This flag just freezes the time by stopping the Renderer calling updateTime()
//
func (w *WidgetClock) SetRunning(b bool) {
	w.running = b
}

//
// Returns true if the clock is running.
//
func (w *WidgetClock) GetRunning() bool {
	return w.running
}

//
// The minimum size of the widget.
// Set by the W and H parameters when created
//
func (w *WidgetClock) MinSize() fyne.Size {
	return w.minSize
}

//
// Pass Resize events to the Base Widget
//
func (w *WidgetClock) Resize(s fyne.Size) {
	w.BaseWidget.Resize(s)
}

func (w *WidgetClock) TappedSecondary(pe *fyne.PointEvent) {
	fmt.Println("TappedSecondary")
}

func (w *WidgetClock) Tapped(pe *fyne.PointEvent) {
	fmt.Println("Tapped")
}
