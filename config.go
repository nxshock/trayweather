package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ncruces/zenity"
)

type Config struct {
	CityName        string `toml:"CityName" env:"CITY_NAME"`
	UpdatePeriodMin int    `toml:"UpdatePeriod", env:"UPDATE_PERIOD"`
}

var config Config

func init() {
	log.SetFlags(0)

	err := cleanenv.ReadConfig("config.toml", &config)
	if err != nil {
		zenity.Notify("Ошибка при чтении настроек из файла config.toml:\n" + err.Error())
		log.Fatalln(err)
	}

	if config.CityName == "" {
		zenity.Notify("Город (поле CityName) не может быть пустым.")
		log.Fatalln("Город (поле CityName) не может быть пустым.")
	}

	if config.UpdatePeriodMin <= 0 {
		log.Printf("Частота обновлений слишком низкая (%s), будет установлено значение в одну минуту.")
		config.UpdatePeriodMin = 1
	}
}
