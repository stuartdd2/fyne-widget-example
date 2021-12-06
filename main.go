package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	ShowWidgetTutorial()
	//	ShowWidgetClockInContainer()
}
func ShowWidgetTutorial() {
	app := app.New()
	w := app.NewWindow("My Widget")
	mw := NewMyWidget("Widget")
	w.SetContent(mw)
	w.ShowAndRun()
}

func ShowWidgetClockInContainer() {
	app := app.New()
	aWindow := app.NewWindow("Widget Clock")
	clock := NewWidgetClock(100, 100, 10, canvas.NewImageFromFile("clock2.svg"), true)
	aButton := widget.NewButton("Stop/Start", func() {
		clock.SetRunning(!clock.GetRunning())
	})
	aWindow.SetContent(container.NewBorder(aButton, nil, nil, nil, clock))
	aWindow.ShowAndRun()
}

func ShowWidgetClock() {
	app := app.New()
	aWindow := app.NewWindow("Widget Clock")
	clock := NewWidgetClock(100, 100, 10, canvas.NewImageFromFile("clock2.svg"), true)
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
