package gui

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/gabrielmoura/FyneT/mqtt"
	"image/color"
	"log"
	"strconv"
)

type App struct {
	A fyne.App
	W fyne.Window
	M mqtt.Mqtt
}
type FlowMeter struct {
	Flowrate  string `json:"flowrate"`
	Flowspeed string `json:"flowspeed"`
	Totalizer string `json:"totalizer"`
	Name      string `json:"NAME"`
	Devid     string `json:"DEVID"`
	Time      any    `json:"TIME"`
}

func (app *App) WPrincipal() {
	app.W = app.A.NewWindow("Vazão da Água")

	copyR := canvas.NewText("Todos os Direitos Reservados", color.RGBA{R: 153, G: 204, B: 50})

	display := widget.NewLabel("Bem Vindo")
	bomba := false

	strFlowrate := binding.NewString()
	strFlowSpeed := binding.NewString()
	app.M.Sub("/flowmeter/loa_duri", func(msg string) {
		testar := FlowMeter{}
		err := json.Unmarshal([]byte(msg), &testar)
		if err != nil {
			log.Fatal(err)
		}
		display.SetText("Vazão da Água")
		strFlowrate.Set(fmt.Sprintf("Flowrate: %s", testar.Flowrate))
		strFlowSpeed.Set(fmt.Sprintf("FlowSpeed: %s", testar.Flowspeed))
	})
	textFlowrate := widget.NewLabelWithData(strFlowrate)
	textFlowSpeed := widget.NewLabelWithData(strFlowSpeed)

	pupup := widget.NewPopUp(container.NewCenter(widget.NewLabel("Exemplo POPUP")), app.W.Canvas())

	home := container.NewVBox(
		container.NewCenter(display),
		container.NewVBox(
			container.NewCenter(textFlowrate),
			container.NewCenter(textFlowSpeed),
		),
		widget.NewButton("Bomba", func() {
			bomba = !bomba
			popupBomba := "Bomba Acionada"
			if !bomba {
				popupBomba = "Bomba Acionada"
			} else {
				popupBomba = "Bomba Desligada"
			}

			app.M.Pub("/bomba", strconv.FormatBool(bomba))
			widget.NewPopUp(container.NewCenter(widget.NewLabel(popupBomba)), app.W.Canvas()).Show()

		}),
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
	//otherItem := fyne.NewMenuItem("Other", nil)
	//otherItem.ChildMenu = fyne.NewMenu("",
	//	fyne.NewMenuItem("Project", func() { fmt.Println("Menu New->Other->Project") }),
	//	fyne.NewMenuItem("Mail", func() { fmt.Println("Menu New->Other->Mail") }),
	//)
	newItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("File", func() { fmt.Println("Menu New->File") }),
		fyne.NewMenuItem("Directory", func() { fmt.Println("Menu New->Directory") }),
		//otherItem,
	)
	layoutItem := fyne.NewMenuItem("Configurações", func() {
		w := app.A.NewWindow("Configurações de Layout")
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
	app.W.SetMainMenu(mainMenu)
	app.W.SetMaster()

	app.W.SetContent(home)

	app.W.Show()

}
func (app *App) Run() {
	app.A.Run()
}
