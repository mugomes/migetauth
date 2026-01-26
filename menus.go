// Copyright (C) 2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package main

import (
	"net/url"

	"fyne.io/fyne/v2"
)

func MainMenus(app fyne.App) *fyne.MainMenu {
	mnuConta := fyne.NewMenu("Conta",
		fyne.NewMenuItem("Gerenciar", func() {

		}),
	)

	mnuAbout := fyne.NewMenu("Sobre",
		fyne.NewMenuItem("Verificar Atualização", func() {
			sURL, _ := url.Parse("https://github.com/mugomes/migetauth/releases")
			app.OpenURL(sURL)

		}),
		fyne.NewMenuItem("Apoie MiGetAuth", func() {
			sURL, _ := url.Parse("https://mugomes.github.io/apoie.html")
			app.OpenURL(sURL)
		}),
		fyne.NewMenuItem("Sobre", func() {
			showAbout(app)
		}),
	)

	return fyne.NewMainMenu(mnuConta, mnuAbout)
}
