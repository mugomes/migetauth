// Copyright (C) 2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package mgprogressbar

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type MGProgressBar struct {
	widget.BaseWidget

	Min   float64
	Max   float64
	Value float64

	height float32

	bgColor color.Color
	fgColor color.Color
}

func New() *MGProgressBar {
	p := &MGProgressBar{
		Min:    0,
		Max:    100,
		Value:  0,
		height: 7,

		bgColor: color.RGBA{200, 200, 200, 255},
		fgColor: color.RGBA{70, 130, 255, 255},
	}
	p.ExtendBaseWidget(p)
	return p
}

type MGProgressBarRenderer struct {
	bar     *MGProgressBar
	bgRect  *canvas.Rectangle
	fgRect  *canvas.Rectangle
	objects []fyne.CanvasObject
}

func (p *MGProgressBar) CreateRenderer() fyne.WidgetRenderer {
	bg := canvas.NewRectangle(p.bgColor)
	fg := canvas.NewRectangle(p.fgColor)

	r := &MGProgressBarRenderer{
		bar:    p,
		bgRect: bg,
		fgRect: fg,
		objects: []fyne.CanvasObject{
			bg,
			fg,
		},
	}
	r.Refresh()
	return r
}

func (r *MGProgressBarRenderer) Layout(size fyne.Size) {
	r.bgRect.Resize(size)

	percent := 0.0
	if r.bar.Max > r.bar.Min {
		percent = (r.bar.Value - r.bar.Min) / (r.bar.Max - r.bar.Min)
	}
	percent = math.Max(0, math.Min(1, percent))

	fgWidth := float32(percent) * size.Width
	r.fgRect.Resize(fyne.NewSize(fgWidth, size.Height))
}

func (r *MGProgressBarRenderer) MinSize() fyne.Size {
	return fyne.NewSize(50, r.bar.height)
}

func (r *MGProgressBarRenderer) Refresh() {
	r.bgRect.FillColor = r.bar.bgColor
	r.fgRect.FillColor = r.bar.fgColor

	r.Layout(r.bar.Size())

	canvas.Refresh(r.bgRect)
	canvas.Refresh(r.fgRect)
}

func (r *MGProgressBarRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *MGProgressBarRenderer) Destroy() {}

func (p *MGProgressBar) SetValue(v float64) {
	p.Value = v
	p.Refresh()
}

func (p *MGProgressBar) SetMin(value float64) {
	p.Min = value
}

func (p *MGProgressBar) SetMax(value float64) {
	p.Max = value
}

func (p *MGProgressBar) SetColors(bg, fg color.Color) {
	p.bgColor = bg
	p.fgColor = fg
}

func (p *MGProgressBar) MinSize() fyne.Size {
	return fyne.NewSize(50, p.height)
}
