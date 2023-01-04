package main

import (
	"Airport/internal/pkg/config"
	mqttConfig "Airport/internal/pkg/mqtt"
	"encoding/json"
	"fmt"
	"math"
	_ "math"
	"strconv"
	"time"
)

const (
	// Constantes
	T0 = 288.15  // Température de l'air à niveau de la mer, en kelvins
	L  = -0.0065 // Lapse rate, en kelvins/mètre
)

// Fonction qui calcule le nombre de minutes écoulées depuis le début de la journée
func elapsedMinutes() float64 {
	now := time.Now()
	return float64(now.Hour()*60 + now.Minute())
}

// Fonction qui calcule la température de l'air en fonction de l'altitude, du temps écoulé et de la saison
func temperature(altitude, time float64, season string) float64 {
	// Ajout d'une correction en fonction de la saison
	correction := 0.0
	if season == "été" {
		correction = 5.0
	} else if season == "hiver" {
		correction = -5.0
	}

	return T0 + L*altitude - L*time + correction
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
	if err == nil {
		fmt.Println("Tempereture sensor")
		client := mqttConfig.Connect(urlBroker, sensor.Id) //TODO
		// Infinit loop for publish each "frenquency" secondes
		for {
			// Calcul du temps écoulé depuis le début de la journée, en minutes
			elapsedTime := elapsedMinutes()

			// Récupération de la date et de la saison actuelles
			now := time.Now()
			season := findSeason(now)
			// Calcul de la température en fonction de l'altitude, du temps écoulé et de la saison
			temp := temperature(float64(altitude), elapsedTime, season)

			tempC := temp - 273.5

			msg := mqttConfig.MessageSensorPublisher{
				SensorId:      sensorId,
				SensorType:    "temperature",
				AirportCode:   sensor.Airport,
				Timestamp:     time.Now(),
				Value:         math.Round(tempC*100) / 100,
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
			fmt.Printf("La température à %v mètres d'altitude est de %.2f °C.\n", altitude, tempC)

			// Attente de 10 secondes
			time.Sleep(time.Duration(frequency) * time.Second)
		}
	}
}
