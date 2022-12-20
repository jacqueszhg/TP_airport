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

var SAISON = map[string]float64{
	"SUMMER": 30.0,
	"SPRING": 20.0,
	"AUTUMN": 10.0,
	"WINTER": 0.0,
}

var LOCATION = map[string]float64{
	"EQUATOR":     10.0,
	"CONTINENTAL": 5.0,
	"NORTH_POLE":  2.0,
	"SOUTH_POLE":  2.0,
}

func simulateTemp(saison float64, location float64, time time.Time, oldLocation float64, oldTime float64) (float64, float64, float64) {
	locationTemp := oldLocation + rand.Float64()*(location-oldLocation)
	timeTemp := oldTime + rand.Float64()*((math.Sin(float64(time.Hour())/math.Pi)*location)-oldTime)
	return saison + locationTemp - timeTemp, locationTemp, timeTemp
}

func main() {
	configPub := config.GetSensorConfig("./config.yml")
	sensor := configPub.Sensor
	mqtt := configPub.MQTT

	urlBroker := mqtt.Protocol + "://" + mqtt.Url + ":" + mqtt.Port
	sensorId, err := strconv.Atoi(sensor.Id)
	QOSLevel, err := strconv.Atoi(sensor.QOSLevel)
	frequency, err := strconv.Atoi(sensor.Frequency)

	if err == nil {
		fmt.Println(urlBroker)

		client := mqttConfig.Connect(urlBroker, sensor.Id) //TODO
		currentTime := time.Now()
		oldLocation := 0.0
		oldTime := 0.0
		currentTemp := 0.0
		// Infinit loop for publish each "frenquency" secondes
		for {
			currentTemp, oldLocation, oldTime = simulateTemp(SAISON["SUMMER"], LOCATION["EQUATOR"], currentTime, oldLocation, oldTime)
			msg := mqttConfig.MessageSensorPublisher{
				SensorId:    sensorId,
				SensorType:  "temperature",
				AirportCode: sensor.Airport,
				Timestamp:   time.Now().String(),
				Value:       currentTemp,
			}

			bytesMsg, err := json.Marshal(msg)

			if err != nil {
				fmt.Println("Can't serislize", msg)
			}
			token := client.Publish("airport/temperature", byte(QOSLevel), true, bytesMsg)
			token.Wait()
			time.Sleep(time.Duration(frequency) * time.Second)
			currentTime.Add(time.Duration(frequency) * (time.Second * 3600))
		}
	}
}
