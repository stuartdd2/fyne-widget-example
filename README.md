# fyne-widget-example
Examples of coding fyne widgets in go

### Rational

While I could find a clock example they did NOT show how a Widget could be constructed from scratch. 

I wanted to explore the proper way to write a Widget in fyne and could not find any **complete** examples.

So I started from scratch and this is the result. 

The following examples are included.

Each example shows the separation of concern for data about a widget (state) and the drawing (rendering) of that widget. There are examples with animators and with timers.

Each example defines a renderer without keeping a reference to that renderer. 

Each example comes with example code that creates, configures and displays the widget.

Each example if fully commented in the source code to help understand why something is the way it is.

Use them as they are or change them as you require. There are NO restrictions on the use of the code.

## WidgetClock

This widget displays a variable sized widget of an analog clock. The clock face is derived using an image from an svg or png image format file. 

The Hour, Minute and Seconds clock hands are drawn by the widget and updated every second.

The colour of the hands can be changed as required. Note the black border is NOT part of the widget.

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
