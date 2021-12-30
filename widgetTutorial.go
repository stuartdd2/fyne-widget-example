package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//
// Widget code starts here
//
// A text widget with theamed background and foreground
//
type MyWidget struct {
	widget.BaseWidget        // Inherit from BaseWidget
	text              string // The text to display in the widget
}

//
// Create a Widget and Extend (initialiase) the BaseWidget
//
func NewMyWidget(text string) *MyWidget {
	w := &MyWidget{ // Create this widget with an initial text value
		text: text,
	}
	w.ExtendBaseWidget(w) // Initialiase the BaseWidget
	return w
}

//
// Create the renderer. This is called by the fyne application
//
func (w *MyWidget) CreateRenderer() fyne.WidgetRenderer {
	// Pass this widget to the renderer so it can access the text field
	return newMyWidgetRenderer(w)
}

//
// Widget Renderer code starts here
//
type myWidgetRenderer struct {
	widget     *MyWidget         // Reference to the widget holding the current state
	background *canvas.Rectangle // A background rectangle
	text       *canvas.Text      // The text
}

//
// Create the renderer with a reference to the widget
// Note: The background and foreground colours are set from the current theme.
//
// Do not size or move canvas objects here.
//
func newMyWidgetRenderer(myWidget *MyWidget) *myWidgetRenderer {
	return &myWidgetRenderer{
		widget:     myWidget,
		background: canvas.NewRectangle(theme.BackgroundColor()),
		text:       canvas.NewText(myWidget.text, theme.ForegroundColor()),
	}
}

//
// The Refresh() method is called if the state of the widget changes or the
// theme is changed
//
// Note: The background and foreground colours are set from the current theme
//
func (r *myWidgetRenderer) Refresh() {
	r.text.Text = r.widget.text
	r.text.Color = theme.ForegroundColor()
	r.background.FillColor = theme.BackgroundColor()
	r.background.Refresh() // Redraw the background first
	r.text.Refresh()       // Redraw the text on top
}

//
// Given the size required by the fyne application move and re-size the
// canvas objects.
//
func (r *myWidgetRenderer) Layout(s fyne.Size) {
	// Measure the size of the text so we can calculate the center offset.
	ts := fyne.MeasureText(r.text.Text, r.text.TextSize, r.text.TextStyle)
	// Center the text
	r.text.Move(fyne.Position{X: (s.Width - ts.Width) / 2, Y: (s.Height - ts.Height) / 2})
	// Make sure the border fills the widget
	r.background.Resize(s)
}

//
// Create a minimum size for the widget.
// The smallest size is the size of the text with a border defined by the theme padding
//
func (r *myWidgetRenderer) MinSize() fyne.Size {
	// Measure the size of the text so we can calculate a border size.
	ts := fyne.MeasureText(r.text.Text, r.text.TextSize, r.text.TextStyle)
	// Use the theme padding to set a border size
	return fyne.NewSize(ts.Width+theme.Padding()*4, ts.Height+theme.Padding()*4)
}

//
// Return a list of each canvas object.
//
func (r *myWidgetRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.background, r.text}
}

//
// Cleanup if resources have been allocated
//
func (r *myWidgetRenderer) Destroy() {}
