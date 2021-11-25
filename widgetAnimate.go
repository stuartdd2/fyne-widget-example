package main

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

//

var _ fyne.WidgetRenderer = (*widgetAnimateRenderer)(nil)
var _ fyne.CanvasObject = (*WidgetAnimate)(nil)
var _ fyne.Widget = (*WidgetAnimate)(nil)
var _ fyne.Tappable = (*WidgetAnimate)(nil)
var _ fyne.SecondaryTappable = (*WidgetAnimate)(nil)
var _ fyne.Disableable = (*WidgetAnimate)(nil)

type widgetAnimator struct {
	animator       *fyne.Animation
	animDelta      float32
	animDeltaReset float32
	animDone       bool
	animObj        fyne.CanvasObject
}

func newWidgetAnimator(min fyne.Size, animObj fyne.CanvasObject, delta float32) *widgetAnimator {
	an := &widgetAnimator{
		animator:       &fyne.Animation{Duration: canvas.DurationStandard},
		animObj:        animObj,
		animDone:       true,
		animDeltaReset: delta,
		animDelta:      delta,
	}
	an.animator.Tick = an.animTick
	return an
}

func (an *widgetAnimator) animTick(f float32) {
	ff := math.Round(float64(f)*10.0) / 10
	if ff == 0.5 {
		an.animDelta = an.animDelta * -1
	}
	if ff == 1.0 {
		an.animator.Stop()
		an.animDone = true
	}
	ns := fyne.NewSize(an.animObj.Size().Width-an.animDelta, an.animObj.Size().Height-an.animDelta)
	if ns.Width > 0 && ns.Height > 0 {
		an.animObj.Resize(ns)
		an.animObj.Move(fyne.NewPos(an.animObj.Position().X+an.animDelta/2, an.animObj.Position().Y+an.animDelta/2))
		an.animObj.Refresh()
	}
}

func (an *widgetAnimator) animate() {
	if an.animDone {
		an.animDone = false
		an.animDelta = an.animDeltaReset
		an.animator.Start()
	}
}

type widgetAnimateRenderer struct {
	minSize fyne.Size
	size    fyne.Size
	circ1   *canvas.Circle
	rect1   *canvas.Rectangle
}

func newWidgetAnimateRenderer(min fyne.Size) (*widgetAnimateRenderer, *widgetAnimator) {
	r := &widgetAnimateRenderer{
		minSize: min, size: min,
		circ1: &canvas.Circle{
			StrokeWidth: 4,
			StrokeColor: color.White,
			FillColor:   color.Transparent,
		}, rect1: &canvas.Rectangle{
			StrokeWidth: 4,
			StrokeColor: color.White,
			FillColor:   color.Transparent,
		},
	}
	return r, newWidgetAnimator(min, r.circ1, 4)
}

func (r *widgetAnimateRenderer) Layout(s fyne.Size) {
	if s.Width <= 1 || s.Height <= 1 {
		return
	}
	r.size = s
	r.circ1.Resize(s)
	r.rect1.Resize(s)
}

func (r *widgetAnimateRenderer) MinSize() fyne.Size {
	return r.minSize
}

func (r *widgetAnimateRenderer) Refresh() {
}

func (r *widgetAnimateRenderer) Objects() []fyne.CanvasObject {
	o := make([]fyne.CanvasObject, 0)
	o = append(o, r.rect1, r.circ1)
	return o
}

func (r *widgetAnimateRenderer) Destroy() {
}

type WidgetAnimate struct {
	widget.BaseWidget
	minSize  fyne.Size
	animator *widgetAnimator
	disabled bool
	tapped   func(*fyne.PointEvent)
}

func NewWidgetAnimate(W, H float32, tapped func(*fyne.PointEvent)) *WidgetAnimate {
	w := &WidgetAnimate{
		disabled: false,
		tapped:   tapped,
		minSize:  fyne.Size{Width: W, Height: H},
	}
	w.ExtendBaseWidget(w)
	return w
}

func (w *WidgetAnimate) TappedSecondary(pe *fyne.PointEvent) {
	if w.tapped != nil && !w.disabled {
		if w.animator != nil {
			w.animator.animate()
		}
		w.tapped(pe)
	}
}

func (w *WidgetAnimate) Tapped(pe *fyne.PointEvent) {
	if w.tapped != nil && !w.disabled {
		if w.animator != nil {
			w.animator.animate()
		}
		w.tapped(pe)
	}
}

func (w *WidgetAnimate) Enable()        { w.disabled = false }
func (w *WidgetAnimate) Disable()       { w.disabled = true }
func (w *WidgetAnimate) Disabled() bool { return w.disabled }

func (w *WidgetAnimate) CreateRenderer() fyne.WidgetRenderer {
	r, a := newWidgetAnimateRenderer(w.MinSize())
	w.animator = a
	return r
}

func (w *WidgetAnimate) MinSize() fyne.Size {
	return w.minSize
}

func (w *WidgetAnimate) Resize(s fyne.Size) {
	w.BaseWidget.Resize(s)
}
