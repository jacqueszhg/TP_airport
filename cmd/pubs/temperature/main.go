package main

import (
	"Airport/internal/pkg/config"
	mqttConfig "Airport/internal/pkg/mqtt"
	"encoding/json"
	"fmt"
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
		msg := mqttConfig.MessageSensorPublisher{
			SensorId:    sensorId,
			SensorType:  "temperature",
			AirportCode: sensor.Airport,
			Timestamp:   time.Now().String(),
			Value:       12.1,
		}

		bytesMsg, err := json.Marshal(msg)

		if err != nil {
			fmt.Println("Can't serislize", msg)
		}

		client := mqttConfig.Connect(urlBroker, sensor.Id) //TODO

		// Infinit loop for publish each "frenquency" secondes
		for {
			token := client.Publish("airport/temperature", byte(QOSLevel), true, bytesMsg)
			token.Wait()
			time.Sleep(time.Duration(frequency) * time.Second)
		}
	}

	//TODO finir le capteur température
}
