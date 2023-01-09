package main

import (
	"Airport/internal/pkg/config"
	mqttConfig "Airport/internal/pkg/mqtt"
	"encoding/json"
	"fmt"
	"math"
	_ "math"
	"math/rand"
	"strconv"
	"time"
)

const (
	// Constantes
	T0 = 288.15  // Température de l'air à niveau de la mer, en kelvins
	L  = -0.0065 // Lapse rate, en kelvins/mètre
)

// Fonction qui calcule la température de l'air en fonction de l'altitude, du temps écoulé et de la saison
func temperature(altitude float64, timeT time.Time, season string) float64 {
	// Ajout d'une correction en fonction de la saison
	correctionSeason := 0.0
	if season == "été" {
		correctionSeason = 5.0
	} else if season == "hiver" {
		correctionSeason = -5.0
	}
	groundTemp := groundTemperature(timeT, season)
	tempWithAltitude := temperatureWithAltitude(altitude)

	// Pour voir des légères variations
	rand.Seed(time.Now().UnixNano())
	randomRange := rand.Float64()
	return tempWithAltitude + groundTemp + correctionSeason + randomRange
}

func temperatureWithAltitude(altitude float64) float64 {
	return T0 + L*altitude
}

func groundTemperature(time time.Time, season string) float64 {
	coef := 3.0 // écart de température air/sol en moyenne
	if season == "été" {
		coef = 5.0
	} else if season == "hiver" {
		coef = 2.0
	}
	return math.Sin(float64(time.Hour()+(time.Minute()/60)+(time.Second()/3600))*math.Pi/24) * coef
}

// Fonction qui renvoie la saison de la date donnée
func findSeason(date time.Time) string {
	year, month, day := date.Date()

	// Calcul du jour de l'année
	n := float64(time.Date(year, month, day, 0, 0, 0, 0, time.UTC).YearDay())

	// Calcul de la saison
	if month == time.December || month == time.January || month == time.February {
		if n >= 355 || n < 79 {
			return "hiver"
		}
	} else if month == time.March || month == time.April || month == time.May {
		if n >= 79 && n < 171 {
			return "printemps"
		}
	} else if month == time.June || month == time.July || month == time.August {
		if n >= 171 && n < 264 {
			return "été"
		}
	} else if month == time.September || month == time.October || month == time.November {
		if n >= 264 && n < 355 {
			return "automne"
		}
	}

	// Saison inconnue
	return "inconnu"
}

func main() {
	configPub := config.GetSensorConfig("./config.yml")
	sensor := configPub.Sensor
	mqtt := configPub.MQTT

	urlBroker := mqtt.Protocol + "://" + mqtt.Url + ":" + mqtt.Port
	sensorId, err := strconv.Atoi(sensor.Id)
	QOSLevel, err := strconv.Atoi(sensor.QOSLevel)
	frequency, err := strconv.Atoi(sensor.Frequency)
	altitude, err := strconv.Atoi(sensor.AltitudeAirport)
	now := time.Now()
	if err == nil {
		fmt.Println("Tempereture sensor")
		client := mqttConfig.Connect(urlBroker, sensor.Id) //TODO
		// Infinit loop for publish each "frenquency" secondes
		for {

			// Récupération de la saison actuelles
			season := findSeason(now)
			// Calcul de la température en fonction de l'altitude, du temps écoulé et de la saison
			temp := temperature(float64(altitude), now, season)

			tempC := temp - 273.5

			msg := mqttConfig.MessageSensorPublisher{
				SensorId:      sensorId,
				SensorType:    "temperature",
				AirportCode:   sensor.Airport,
				Timestamp:     time.Now(),
				Value:         tempC,
				UnitOfMeasure: "Celsius",
			}

			bytesMsg, err := json.Marshal(msg)

			if err != nil {
				fmt.Println("Can't serialize", msg)
			}
			tokenDB := client.Publish("airport/temperature", byte(QOSLevel), true, bytesMsg)
			tokenLog := client.Publish("airport/log", byte(QOSLevel), true, bytesMsg)
			tokenDB.Wait()
			tokenLog.Wait()

			// Affichage de la température
			fmt.Printf("[%dH] La température à %v mètres d'altitude est de %f °C.\n", now.Hour(), altitude, tempC)

			// Attente de 10 secondes
			time.Sleep(time.Duration(frequency) * time.Second)
			now = now.Add(time.Duration(frequency) * time.Second)
		}
	}
}
