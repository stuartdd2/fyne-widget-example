# fyne-widget-example
Examples of coding fyne widgets in go

## Rational
=======

While I could find clock examples they did NOT demonstrate how a Widget could be constructed from scratch. 

I wanted to explore the proper way to write a Widget in fyne and could not find any **complete** examples.

So I started from scratch and this is the result. 

The following examples are included:

* WidgetClock
* WidgetAnimator

Each example shows the separation of concern for data about a widget (state) and the drawing (rendering) of that widget. 

There are examples using both fyne.Animation and time.Ticker.

Each example defines a renderer without keeping a reference to that renderer. 

Each example comes with example code (main.go) that creates, configures and displays the widget.

Each example if fully commented in the source code to help understand why something is the way it is.

Use them as they are or change them as you require. There are NO restrictions on the use of the code.

**Please feed back issues and improvements. I am learning both Go and fyne so there will be improvements that can be made.**

## WidgetClock

This widget displays a variable sized widget of an analog clock. The clock face is derived using an image from an svg or png image format file. 

The Hour, Minute and Seconds clock hands are drawn by the widget and updated every second using time.Ticker.

The colour of the hands can be changed as required. Note the black border is NOT part of the widget in the image below.

![fyne-clock-widget](https://user-images.githubusercontent.com/94919638/143290347-2a0f5f1c-7015-4a00-b994-72b17e272ee1.png)

```go
func main() {
	app := app.New()
	aWindow := app.NewWindow("Widget Clock")
	clock := NewWidgetClock(10, canvas.NewImageFromFile("clock2.svg"), 100, 100, true)
	aWindow.SetContent(clock)
	aWindow.ShowAndRun()
}
```

See source code comments for details.

## WidgetAnimator

This widget displays a simple circle in a rectangle. When the widget is clicked the circle animates using fyne.Animation. It shrinks and then expands to its original size.

This is an example of the use of the fyne.Animation and a fyne.WidgetRenderer combined to animate a simple button like widget and can be used as a starting point for more creative widgets of your own.

![Screenshot from 2021-11-25 17-50-39](https://user-images.githubusercontent.com/94919638/143485360-e7962123-d8f8-4b2c-9ab0-62f1aefbb9cd.png)

See source code comments for details.

