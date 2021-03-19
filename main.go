package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getlantern/systray"
	"github.com/nxshock/trayweather/yandex"
)

//go:embed icons/*.ico
var icons embed.FS

func init() {
	log.SetFlags(0)

	http.DefaultClient.Timeout = 10 * time.Second
}

func main() {
	systray.Run(onReady, nil)
}

func onReady() {
	setTrayIcon("unknown.ico")
	go update()

	mQuit := systray.AddMenuItem("Выход", "Выйти из приложения")
	exitIcon, err := icons.ReadFile("icons/exit.ico")
	if err != nil {
		log.Fatalln(err)
	}
	mQuit.SetIcon(exitIcon)

	go func() {
		<-mQuit.ClickedCh
		os.Exit(0)
	}()
}

func update() {
	for {
		c, err := yandex.Get(config.CityName)
		if err != nil {
			systray.SetTooltip(err.Error())
			setTrayIcon("unknown")
			time.Sleep(time.Minute)
			continue
		}

		systray.SetTooltip(fmt.Sprintf("%s\n%.1f °C (%.1f °C)", c.Description(), c.CurrentTemperature(), c.FeelsLikeTemperature()))
		setTrayIcon(c.IconName())
		time.Sleep(config.UpdatePeriod)
	}
}

func setTrayIcon(name string) error {
	b, err := icons.ReadFile("icons/" + name)
	if err != nil {
		return err
	}
	systray.SetIcon(b)

	return nil
}
