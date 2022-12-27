package main

import (
	"Airport/internal/pkg/config"
	mqttConfig "Airport/internal/pkg/mqtt"
	"encoding/json"
	"fmt"
	_ "math"
	"strconv"
	"time"
)

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
		// Infinit loop for publish each "frenquency" secondes
		for {
			msg := mqttConfig.MessageSensorPublisher{
				SensorId:      sensorId,
				SensorType:    "pressure",
				AirportCode:   sensor.Airport,
				Timestamp:     time.Now(),
				Value:         0.0,
				UnitOfMeasure: "km/h",
			}

			bytesMsg, err := json.Marshal(msg)

			if err != nil {
				fmt.Println("Can't serislize", msg)
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
