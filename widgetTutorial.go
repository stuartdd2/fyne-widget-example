package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type myWidgetRenderer struct {
	widget *MyWidget
	rect   *canvas.Rectangle
	text   *canvas.Text
}

var _ fyne.WidgetRenderer = (*myWidgetRenderer)(nil) // Test myWidgetRenderer is a fyne.fyne.WidgetRenderer
var _ fyne.Tappable = (*MyWidget)(nil)               // Test myWidgetRenderer is a fyne.fyne.WidgetRenderer

func newMyWidgetRenderer(myWidget *MyWidget) *myWidgetRenderer {
	return &myWidgetRenderer{
		widget: myWidget,
		rect:   canvas.NewRectangle(theme.BackgroundColor()),
		text:   canvas.NewText(myWidget.text, theme.ForegroundColor()),
	}
}

func (r *myWidgetRenderer) Refresh() {
	r.text.Text = r.widget.text
}

func (r *myWidgetRenderer) Layout(s fyne.Size) {
	ts := fyne.MeasureText(r.text.Text, r.text.TextSize, r.text.TextStyle)
	r.text.Move(fyne.Position{X: (s.Width - ts.Width) / 2, Y: (s.Height - ts.Height) / 2})
	r.rect.Resize(s)
}

func (r *myWidgetRenderer) MinSize() fyne.Size {
	ts := fyne.MeasureText(r.text.Text, r.text.TextSize, r.text.TextStyle)
	return fyne.NewSize(ts.Width+20, ts.Height+20)
}

func (r *myWidgetRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rect, r.text}
}

func (r *myWidgetRenderer) Destroy() {}

type MyWidget struct {
	widget.BaseWidget        // Inherit from BaseWidget
	text              string // The text to display in the widget
}

//
// Create a Widget and Extend the BaseWidget
//
func NewMyWidget(text string) *MyWidget {
	w := &MyWidget{
		text: text,
	}
	w.ExtendBaseWidget(w)
	return w
}

func (w *MyWidget) CreateRenderer() fyne.WidgetRenderer {
	return newMyWidgetRenderer(w)
}

func (w *MyWidget) Tapped(*fyne.PointEvent) {
	fmt.Println("MyWidget h as been tapped")
	w.text = "FRED"
	w.Refresh()
}
