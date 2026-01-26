// Copyright (C) 2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package main

import (
	"database/sql"
	"log"
	"mugomes/migetauth/controls"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/mugomes/mgsmartflow"
)

func (config *MiGetAuthConfig) showCreateAccount(a fyne.App, db *sql.DB) {
	window := a.NewWindow("MiGetAuth")
	window.SetFixedSize(true)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(800, 600))

	flow := mgsmartflow.New()

	lblNome := widget.NewLabel("Nome/Apelido")
	lblNome.TextStyle = fyne.TextStyle{Bold: true}
	txtNome := widget.NewEntry()

	flow.AddRow(lblNome)
	flow.AddRow(txtNome)

	lblEmail := widget.NewLabel("E-mail")
	lblEmail.TextStyle = fyne.TextStyle{Bold: true}
	txtEmail := widget.NewEntry()

	flow.AddRow(lblEmail)
	flow.AddRow(txtEmail)

	lblSenha := widget.NewLabel("Senha")
	lblSenha.TextStyle = fyne.TextStyle{Bold: true}
	txtSenha := widget.NewPasswordEntry()

	flow.AddRow(lblSenha)
	flow.AddRow(txtSenha)

	btnCreate := widget.NewButton("Criar Conta", func() {
		config.sAccount = txtSenha.Text
		err := controls.CreateUser(db, txtNome.Text, txtEmail.Text, txtSenha.Text)
		if err != nil {
			log.Fatal(err)
		}

		createCode()
		window.Close()
	})

	flow.AddRow(container.NewHBox(layout.NewSpacer(), btnCreate, layout.NewSpacer()))

	window.SetContent(flow.Container)
	window.Show()
}
