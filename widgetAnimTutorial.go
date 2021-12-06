package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type myWidgetAnimRenderer struct {
	minSize fyne.Size // The size of the widget
	rect    *canvas.Rectangle
	text    *canvas.Text
}

var _ fyne.WidgetRenderer = (*myWidgetAnimRenderer)(nil) // Test myWidgetAnimRenderer is a fyne.fyne.WidgetRenderer
var _ fyne.Tappable = (*MyWidgetAnim)(nil)               // Test MyWidgetAnim is a fyne.fyne.Tappable

func newMyWidgetAnimRenderer(size fyne.Size, text string, textSize float32) *myWidgetAnimRenderer {
	return &myWidgetAnimRenderer{
		minSize: size,
		rect:    canvas.NewRectangle(theme.BackgroundColor()),
		text:    &canvas.Text{Text: text, TextSize: textSize},
	}
}

func (r *myWidgetAnimRenderer) Refresh() {}

func (r *myWidgetAnimRenderer) Layout(s fyne.Size) {
	si := fyne.MeasureText(r.text.Text, r.text.TextSize, r.text.TextStyle)
	r.text.Move(fyne.Position{X: (s.Width - si.Width) / 2, Y: (s.Height - si.Height) / 2})
	r.rect.Resize(s)
}

func (r *myWidgetAnimRenderer) MinSize() fyne.Size {
	return r.minSize
}

func (r *myWidgetAnimRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rect, r.text}
}

func (r *myWidgetAnimRenderer) ObjectsForAnimation() *canvas.Text {
	return r.text
}

func (r *myWidgetAnimRenderer) Destroy() {
}

type myWidgetAnimControl struct {
	animator *fyne.Animation // The fyne animator
	text     *canvas.Text
	textSize float32
}

func newWidgetAnimControl(text *canvas.Text) *myWidgetAnimControl {
	anim := &myWidgetAnimControl{
		text:     text,
		textSize: text.TextSize,
		animator: &fyne.Animation{
			Duration: canvas.DurationStandard,
		},
	}
	anim.animator.Tick = anim.animate
	return anim
}

func (a myWidgetAnimControl) animate(f float32) {
	a.text.TextSize = a.textSize * f
	canvas.Refresh(a.text)
}

type MyWidgetAnim struct {
	widget.BaseWidget                      // Inherit from BaseWidget
	minSize           fyne.Size            // The minimum size of the widget
	text              string               // The text to display in the widget
	animator          *myWidgetAnimControl // The fyne animator
}

//
// Create a Widget and Extend the BaseWidget
//
func NewMyWidgetAnim(W, H float32, text string) *MyWidgetAnim {
	w := &MyWidgetAnim{
		text:    text,
		minSize: fyne.Size{Width: W, Height: H},
	}
	w.ExtendBaseWidget(w)
	return w
}

//
// Part of the widget interface. Provided a reference to the Renderer
// Do not hang on to the reference to the Renderer. fyne may stop using
// it and call again for a new one. fyne also caches the reference.
// The Renderer gets the widget state passed in as a reference.
//
func (w *MyWidgetAnim) CreateRenderer() fyne.WidgetRenderer {
	r := newMyWidgetAnimRenderer(w.minSize, w.text, 30)
	w.animator = newWidgetAnimControl(r.ObjectsForAnimation())
	return r
}

//
// MinSize: The minimum size of the widget.
//
func (w *MyWidgetAnim) MinSize() fyne.Size {
	return w.minSize
}

//
// Resize: delegate resize events to the Base Widget
//
func (w *MyWidgetAnim) Resize(s fyne.Size) {
	w.BaseWidget.Resize(s)
}

func (w *MyWidgetAnim) Tapped(*fyne.PointEvent) {
	w.animator.animator.Start()
}
