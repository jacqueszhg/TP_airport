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

func simulatePressure(altitude float64) float64 {

	// Calcule l'heure actuelle en secondes
	heureActuelle := time.Now().Hour()*3600 + time.Now().Minute()*60 + time.Now().Second()

	hauteur := 3.0                //hauteur du capteur par rapport au sol en m
	masseAir := 5.0               //maisse de l'air en kg
	accelerationPesanteur := 9.81 //acceleration pesanteur en m/s^2

	// Calcule la pression atmosphérique en utilisant une sinusoïde et en tenant compte de l'altitude
	pression := math.Sin(float64(heureActuelle)/(12*3600)*math.Pi)*masseAir*accelerationPesanteur*(hauteur+altitude) + 101325 // pression en Pascals

	hPaPression := pression / 100

	return hPaPression
}

func main() {
	configPub := config.GetSensorConfig("./config.yml")
	sensor := configPub.Sensor
	mqtt := configPub.MQTT

	urlBroker := mqtt.Protocol + "://" + mqtt.Url + ":" + mqtt.Port
	sensorId, err := strconv.Atoi(sensor.Id)
	QOSLevel, err := strconv.Atoi(sensor.QOSLevel)
	frequency, err := strconv.Atoi(sensor.Frequency)
	altitudeAirport, err := strconv.Atoi(sensor.AltitudeAirport)

	if err == nil {
		fmt.Println("Pressure sensor")

		client := mqttConfig.Connect(urlBroker, sensor.Id)
		currentTime := time.Now()

		// Infinit loop for publish each "frenquency" secondes
		for {
			currentPressure := simulatePressure(float64(altitudeAirport))

			msg := mqttConfig.MessageSensorPublisher{
				SensorId:      sensorId,
				SensorType:    "pressure",
				AirportCode:   sensor.Airport,
				Timestamp:     time.Now(),
				Value:         currentPressure,
				UnitOfMeasure: "hPa",
			}

			bytesMsg, err := json.Marshal(msg)

			if err != nil {
				fmt.Println("Can't serialize", msg)
			}
			tokenDB := client.Publish("airport/pressure", byte(QOSLevel), true, bytesMsg)
			tokenLog := client.Publish("airport/log", byte(QOSLevel), true, bytesMsg)
			tokenDB.Wait()
			tokenLog.Wait()
			time.Sleep(time.Duration(frequency) * time.Second)
			currentTime.Add(time.Duration(frequency) * (time.Second * 3600))
		}
	}
}
