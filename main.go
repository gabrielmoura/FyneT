package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/gabrielmoura/FyneT/gui"
	"github.com/gabrielmoura/FyneT/mDNS"
	"github.com/gabrielmoura/FyneT/mqtt"
)

func main() {
	name := "br.com.srmoura.demo"
	mDNS.SetDNS(name)
	mm := mqtt.Mqtt{}
	mm.Connect(mqtt.MqttConfig{
		ClientId: name,
		Broker:   "tcp://broker.emqx.io:1883",
		Username: "emqx",
		Password: "public",
	})
	X := gui.App{A: app.NewWithID(name), M: mm}

	X.WPrincipal()
	X.Run()
}
