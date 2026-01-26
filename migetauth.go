// Copyright (C) 2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"

	"github.com/mugomes/mgsmartflow"

	"mugomes/migetauth/components/mgprogressbar"
	c "mugomes/migetauth/controls"
)

const VERSION_APP = "1.0.0"

func getCode() string {
	secret := ""
	code, err := c.GenerateTOTP(secret)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return code
}

func main() {
	app := app.NewWithID("br.com.mugomes.migetauth")
	app.Settings().SetTheme(&myDarkTheme{})
	window := app.NewWindow("MiGetAuth")
	window.SetFixedSize(true)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(800, 600))

	flow := mgsmartflow.New()

	window.SetMainMenu(c.MainMenus(app))

	lblContaCodigo := canvas.NewText(getCode(), color.White)
	lblContaCodigo.TextSize = 32
	lblContaCodigo.Alignment = fyne.TextAlignCenter
	lblContaNome := canvas.NewText("", color.White)
	lblContaNome.TextSize = 18
	lblContaNome.Alignment = fyne.TextAlignCenter

	flow.AddRow(lblContaCodigo)
	flow.AddRow(lblContaNome)

	pTempo := mgprogressbar.New()
	pTempo.Min = 0
	pTempo.Max = 15
	flow.AddRow(pTempo)

	go (func() {
		tckTempo := time.NewTicker(1 * time.Second)
		defer tckTempo.Stop()

		sSeconds := 0

		for range tckTempo.C {
			sSeconds++

			fyne.Do(func() {
				pTempo.SetValue(float64(sSeconds))
			})

			if sSeconds == int(pTempo.Max) {
				sSeconds = 0

				fyne.Do(func() {
					lblContaCodigo.Text = getCode()
					lblContaCodigo.Refresh()
				})
			}
		}
	})()

	window.SetContent(flow.Container)
	window.ShowAndRun()
}
