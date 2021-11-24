package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	ShowWidgetAnimator()
	//	ShowWidgetClockInContainer()
}

func ShowWidgetClockInContainer() {
	app := app.New()
	aWindow := app.NewWindow("Widget Clock")
	clock := NewWidgetClock(10, canvas.NewImageFromFile("clock2.svg"), 100, 100, true)
	aButton := widget.NewButton("Stop/Start", func() {
		clock.SetRunning(!clock.GetRunning())
	})
	aWindow.SetContent(container.NewBorder(aButton, nil, nil, nil, clock))
	aWindow.ShowAndRun()
}
func ShowWidgetClock() {
	app := app.New()
	aWindow := app.NewWindow("Widget Clock")
	clock := NewWidgetClock(10, canvas.NewImageFromFile("clock2.svg"), 100, 100, true)
	aWindow.SetContent(clock)
	aWindow.ShowAndRun()
}

func ShowWidgetAnimator() {
	var wa *WidgetAnimate
	app := app.New()
	aWindow := app.NewWindow("Widget Animator")
	wa = NewWidgetAnimate(50, 50, func(pe *fyne.PointEvent) {})
	aWindow.SetContent(wa)
	aWindow.ShowAndRun()
}
