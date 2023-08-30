package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *App) WSecond() {
	app.W = app.A.NewWindow("Montador2")

	hello := widget.NewLabel("Bem vindo 2")
	matar := widget.NewLabel("Matar")

	app.W.CenterOnScreen()

	app.W.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Montar", func() {
			hello.SetText("Montado :)")
			not := fyne.NewNotification("Test", "Completo")
			app.A.SendNotification(not)

		}),
		widget.NewButtonWithIcon("Fechar", theme.CancelIcon(), func() {
			app.W.Close()
		}),
		matar,
	))
	app.W.Show()
}
