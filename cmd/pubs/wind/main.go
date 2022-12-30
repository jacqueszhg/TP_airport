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

func getWind(min float64, max float64) float64 {
	r := min + rand.Float64()*(max-min)
	return r
}

func main() {

	sensor := config.GetSensorConfig("./config.yml").Sensor
	mqtt := config.GetSensorConfig("./config.yml").MQTT

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
				Value:         getWind(0.4, 80),
				UnitOfMeasure: "m/s",
			}

			jsonMessage, jsonErr := json.Marshal(message)

			if jsonErr == nil {
				tokenDB := client.Publish("airport/temperature", byte(QOS), true, jsonMessage)
				tokenLog := client.Publish("airport/log", byte(QOS), true, jsonMessage)
				tokenDB.Wait()
				tokenLog.Wait()
				time.Sleep(time.Duration(frequency) * time.Second)
			}
		}
	}

}
