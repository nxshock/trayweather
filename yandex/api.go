package yandex

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type WeatherData struct {
	currentTemperature   float64
	feelsLikeTemperature float64
	description          string
	iconURL              string
}

func (w *WeatherData) CurrentTemperature() float64 {
	return w.currentTemperature
}

func (w *WeatherData) FeelsLikeTemperature() float64 {
	return w.feelsLikeTemperature
}
func (w *WeatherData) Description() string {
	return w.description
}
func (w *WeatherData) IconName() string {
	switch w.description {
	case "Ясно", "Малооблачно":
		return "01d.ico"
	case "Облачно с прояснениями":
		return "02d.ico"
	case "Пасмурно":
		return "03d.ico"
	case "Небольшой снег":
		return "09d.ico"
	case "Снег":
		return "13d.ico"
	}
	return ""
}

func Get(cityName string) (*WeatherData, error) {
	url := fmt.Sprintf("https://yandex.ru/pogoda/%s", strings.ToLower(cityName))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	currentTemp, err := parseFloat(doc.Find("div.fact__temp > span.temp__value").Text())
	if err != nil {
		return nil, fmt.Errorf("parse current temperature error: %v", err)
	}

	feelsLikeTemp, err := parseFloat(doc.Find("div.fact__feels-like > div.term__value span.temp__value").Text())
	if err != nil {
		return nil, fmt.Errorf("parse feels like temperature error: %v", err)
	}

	wd := &WeatherData{
		currentTemperature:   currentTemp,
		feelsLikeTemperature: feelsLikeTemp,
		description:          doc.Find("div.link__condition").Text(),
		iconURL:              doc.Find("div.fact__temp-wrap img.fact__icon").AttrOr("src", "")}

	return wd, nil
}

func parseFloat(s string) (float64, error) {
	var b []rune

	s = strings.ReplaceAll(s, ",", ".")
	s = strings.ReplaceAll(s, "−", "-")

	for _, r := range []rune(s) {
		if (r >= '0' && r <= '9') || (r == '+' || r == '-') || (r == '.') {
			b = append(b, r)
		}
	}

	return strconv.ParseFloat(string(b), 32)
}
