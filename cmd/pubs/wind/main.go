package main

import (
	"Airport/internal/pkg/config"
	mqttConfig "Airport/internal/pkg/mqtt"
	"encoding/json"
	_ "math"
	"math/rand"
	"strconv"
	"time"
)

func getWind(lastValue float64) float64 {
	r := -10 + rand.Float64()*(20)
	newValue := lastValue + r
	for r > 0.4 && r < 80 {
		r = -10 + rand.Float64()*(20)
		newValue = lastValue + r
	}
	return newValue
}

func main() {

	configPub := config.GetSensorConfig("./config.yml")
	sensor := configPub.Sensor
	mqtt := configPub.MQTT
	urlBroker := mqtt.Protocol + "://" + mqtt.Url + ":" + mqtt.Port

	sensorId, err := strconv.Atoi(sensor.Id)
	QOS, err := strconv.Atoi(sensor.QOSLevel)
	frequency, err := strconv.Atoi(sensor.Frequency)

	if err == nil {

		client := mqttConfig.Connect(urlBroker, sensor.Id)

		for {
			message := mqttConfig.MessageSensorPublisher{
				SensorId:      sensorId,
				SensorType:    "wind",
				AirportCode:   sensor.Airport,
				Timestamp:     time.Now(),
				Value:         getWind(0.4 + rand.Float64()*(80-0.4)), // modern wind detector can record wind speed between 0.4 m/s and 80 m/s (the global unit)
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
		}
	}

}
