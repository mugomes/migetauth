// Copyright (C) 2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package main

import (
	"database/sql"
	"fmt"
	"image/color"
	"log"
	"os"
	"path"
	"path/filepath"

	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/widget"

	"github.com/mugomes/mgsmartflow"

	"mugomes/migetauth/components/mgprogressbar"
	c "mugomes/migetauth/controls"
)

const VERSION_APP string = "1.0.0"

type MiGetAuthConfig struct {
	dbFileName string
	sIDUser    int64
}

func getCode(secret string) string {
	code, err := c.GenerateTOTP(secret)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return code
}

func initConfig(filename string, iduser int64) *MiGetAuthConfig {
	sDir := path.Dir(filename)
	if !c.FileExists(sDir) {
		os.Mkdir(sDir, os.ModePerm)
	}

	return &MiGetAuthConfig{
		dbFileName: filename,
		sIDUser: iduser,
	}
}

var (
	pTempo         *mgprogressbar.MGProgressBar
	lblContaCodigo *canvas.Text

	db  *sql.DB
	err error
)

func createCode() {
	go (func() {
		accountExists, _ := c.AccountExists(db)
		if accountExists {
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
						lblContaCodigo.Text = getCode("JBSWY3DPEHPK3PXP")
						lblContaCodigo.Refresh()
					})
				}
			}
		}
	})()
}

func main() {
	app := app.NewWithID("br.com.mugomes.migetauth")
	app.Settings().SetTheme(&myDarkTheme{})

	fyne.Do(func() {
		sIcon := fyne.NewStaticResource("migetauth.png", resourceAppIconData)
		app.SetIcon(sIcon)
	})

	window := app.NewWindow("MiGetAuth")
	window.SetFixedSize(true)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(800, 299))

	flow := mgsmartflow.New()

	window.SetMainMenu(MainMenus(app))

	appConfig := initConfig(filepath.Join("data", "migetauth.db"), 0)

	lblContaCodigo = canvas.NewText("", color.White)
	lblContaCodigo.TextSize = 32
	lblContaCodigo.Alignment = fyne.TextAlignCenter
	lblContaNome := canvas.NewText("", color.White)
	lblContaNome.TextSize = 18
	lblContaNome.Alignment = fyne.TextAlignCenter

	flow.AddRow(lblContaCodigo)
	flow.AddRow(lblContaNome)

	pTempo = mgprogressbar.New()
	pTempo.Min = 0
	pTempo.Max = 30
	flow.AddRow(pTempo)

	// widget.NewButtonWithIcon("", "", func() {

	// })

	window.SetContent(flow.Container)

	db, err = sql.Open("sqlite3", appConfig.dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbErr := c.CreateTable(db)
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	dbUserExists, err := c.UserExists(db)
	if err != nil {
		log.Fatal(err)
	}

	if !dbUserExists {
		appConfig.showCreateUser(app, db, window)
	} else {
		appConfig.showUserLogin(app, db, window)
	}

	app.Run()
}
