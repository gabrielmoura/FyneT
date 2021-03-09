package main

import (
	"image/color"

	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	a fyne.App
	w fyne.Window
}

func (app *App) WPrincipal() {
	app.w = app.a.NewWindow("Montador")

	copyR := canvas.NewText("Todos os direitos Reservados", color.RGBA{R: 153, G: 204, B: 50})

	display := widget.NewLabel("Bem vindo")

	pupup := widget.NewPopUp(container.NewCenter(widget.NewLabel("Exemplo POPUP")), app.w.Canvas())

	home := container.NewVBox(
		container.NewCenter(display),
		widget.NewButton("Montar", func() {
			display.SetText("Montado :)")
			app.WSecond()
		}),
		widget.NewButton("Desmontar", func() {
			display.SetText("Desmontado :)")

			pupup.Show()
			//widget.ShowPopUp(popup)
		}),
		container.NewCenter(copyR))

	newItem := fyne.NewMenuItem("New", nil)
	otherItem := fyne.NewMenuItem("Other", nil)
	otherItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Project", func() { fmt.Println("Menu New->Other->Project") }),
		fyne.NewMenuItem("Mail", func() { fmt.Println("Menu New->Other->Mail") }),
	)
	newItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("File", func() { fmt.Println("Menu New->File") }),
		fyne.NewMenuItem("Directory", func() { fmt.Println("Menu New->Directory") }),
		otherItem,
	)
	layoutItem := fyne.NewMenuItem("Configurações", func() {
		w := app.a.NewWindow("Configurações de Layout")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	})

	mainMenu := fyne.NewMainMenu(
		// a quit item will be appended to our first menu
		fyne.NewMenu("Menu", newItem, fyne.NewMenuItemSeparator(), layoutItem),
		//fyne.NewMenu("Edit", cutItem, copyItem, pasteItem, fyne.NewMenuItemSeparator(), findItem),
		//helpMenu,
	)
	app.w.SetMainMenu(mainMenu)
	app.w.SetMaster()

	app.w.SetContent(home)

	app.w.Show()

}

func (app *App) WSecond() {
	app.w = app.a.NewWindow("Montador2")

	hello := widget.NewLabel("Bem vindo 2")
	matar := widget.NewLabel("Matar")

	app.w.CenterOnScreen()

	app.w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Montar", func() {
			hello.SetText("Montado :)")
			not := fyne.NewNotification("Test", "Completo")
			app.a.SendNotification(not)

		}),
		widget.NewButtonWithIcon("Fechar", theme.CancelIcon(), func() {
			app.w.Close()
		}),
		matar,
	))
	app.w.Show()
}

func main() {
	X := App{a: app.NewWithID("br.com.srmoura.demo")}

	X.WPrincipal()
	X.a.Run()
}
