// Copyright (C) 2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package main

import (
	"database/sql"
	"mugomes/migetauth/controls"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/mugomes/mgdialogbox"
	"github.com/mugomes/mgsmartflow"
)

func (config *MiGetAuthConfig) showUserLogin(a fyne.App, db *sql.DB, wMain fyne.Window) {
	window := a.NewWindow("Login")
	window.SetFixedSize(true)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(400, 300))

	flow := mgsmartflow.New()

	lblUsuario := widget.NewLabel("Usuário")
	lblUsuario.TextStyle = fyne.TextStyle{Bold: true}
	txtUsuario := widget.NewEntry()

	flow.AddRow(lblUsuario)
	flow.AddRow(txtUsuario)

	lblSenha := widget.NewLabel("Senha")
	lblSenha.TextStyle = fyne.TextStyle{Bold: true}
	txtSenha := widget.NewPasswordEntry()

	flow.AddRow(lblSenha)
	flow.AddRow(txtSenha)

	flow.Gap(txtSenha, txtSenha.Position().X, 17)

	var sCheckLogin int = 3
	btnAcessar := widget.NewButton("Acessar", func() {
		checkUser, err := controls.CheckUser(db, txtUsuario.Text, txtSenha.Text)
		if err != nil {
			mgdialogbox.NewAlert(a, "MiGetAuth", "Usuário não encontrado!", true, "Continuar")
		}

		if checkUser {
			sCheckLogin = 1
			fyne.Do(func() {
				wMain.Show()
			})

			createCode()
			window.Close()
		} else {
			mgdialogbox.NewAlert(a, "MiGetAuth", "Usuário não encontrado!", true, "Continuar")
		}
	})

	btnCriarConta := widget.NewButton("Criar Conta", func() {
		sCheckLogin = 2
		config.showCreateUser(a, db, wMain)
		window.Close()
	})

	flow.AddRow(container.NewHBox(
		layout.NewSpacer(),
		btnAcessar,
		btnCriarConta,
		layout.NewSpacer(),
	))

	window.SetContent(flow.Container)
	window.Show()

	window.SetOnClosed(func() {
		if sCheckLogin == 1 {
			wMain.Show()
		} else if sCheckLogin > 2 {
			a.Quit()
		}
	})
}
