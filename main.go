package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var aClock *WidgetClock
var aButton *widget.Button

func main() {
	ShowWidgetClock()
}

func ShowWidgetClock() {
	app := app.New()
	aWindow := app.NewWindow("Widget Clock")
	aClock = NewWidgetClock(10, canvas.NewImageFromFile("clock2.svg"), 100, 100)
	aButton = widget.NewButton("Stop/Start", func() {
		aClock.SetRunning(!aClock.GetRunning())
	})

	aWindow.SetContent(container.NewBorder(aButton, nil, nil, nil, aClock))
	aClock.SetRunning(true)
	aWindow.ShowAndRun()
}
