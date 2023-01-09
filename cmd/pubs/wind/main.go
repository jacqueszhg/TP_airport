package main

import (
	"Airport/internal/pkg/config"
	mqttConfig "Airport/internal/pkg/mqtt"
	"encoding/json"
	"fmt"
	_ "math"
	"math/rand"
	"strconv"
	"time"
)

func findSeason(date time.Time) string {
	year, month, day := date.Date()

	// Calcul du jour de l'année
	n := float64(time.Date(year, month, day, 0, 0, 0, 0, time.UTC).YearDay())

	// Calcul de la saison
	if month == time.December || month == time.January || month == time.February {
		if n >= 355 || n < 79 {
			return "winter"
		}
	} else if month == time.March || month == time.April || month == time.May {
		if n >= 79 && n < 171 {
			return "spring"
		}
	} else if month == time.June || month == time.July || month == time.August {
		if n >= 171 && n < 264 {
			return "summer"
		}
	} else if month == time.September || month == time.October || month == time.November {
		if n >= 264 && n < 355 {
			return "autumn"
		}
	}

	// Saison inconnue
	return "inconnu"
}

func getWind(altitude int) float64 {
	season := findSeason(time.Now())
	r := 0.0
	if season == "summer" {
		r = rand.Float64()*(16-9) + 9
	} else if season == "winter" {
		r = rand.Float64()*(9-4) + 4
	} else if season == "spring" {
		r = rand.Float64()*(11-7) + 7
	} else if season == "autumn" {
		r = rand.Float64()*(15-11) + 11
	}

	r += float64(altitude / 500)
	return r
}

func main() {

	configPub := config.GetSensorConfig("./config.yml")
	sensor := configPub.Sensor
	mqtt := configPub.MQTT
	urlBroker := mqtt.Protocol + "://" + mqtt.Url + ":" + mqtt.Port

	sensorId, err := strconv.Atoi(sensor.Id)
	QOS, err := strconv.Atoi(sensor.QOSLevel)
	frequency, err := strconv.Atoi(sensor.Frequency)
	altitude, err := strconv.Atoi(sensor.AltitudeAirport)

	if err == nil {

		client := mqttConfig.Connect(urlBroker, sensor.Id)

		for {
			message := mqttConfig.MessageSensorPublisher{
				SensorId:      sensorId,
				SensorType:    "wind",
				AirportCode:   sensor.Airport,
				Timestamp:     time.Now(),
				Value:         getWind(altitude), // modern wind detector can record wind speed between 0.4 m/s and 80 m/s (the global unit)
				UnitOfMeasure: "m/s",
			}

			jsonMessage, jsonErr := json.Marshal(message)

			if jsonErr == nil {
				tokenDB := client.Publish("airport/wind", byte(QOS), true, jsonMessage)
				tokenLog := client.Publish("airport/log", byte(QOS), true, jsonMessage)
				tokenDB.Wait()
				tokenLog.Wait()
				time.Sleep(time.Duration(frequency) * time.Second)
			}
			fmt.Println(message)
		}
	}

}
