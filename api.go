package main

type WeatherData interface {
	CurrentTemperature() float64
	FeelsLikeTemperature() float64
	Description() string
	IconName() string
}
