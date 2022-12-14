package main

import (
	"Airport/internal/pkg/config"
	mqttConfig "Airport/internal/pkg/mqtt"
	"strconv"
	"time"
)

func main() {
	sensor := config.GetSensorConfig("").Sensor
	mqtt := config.GetSensorConfig("").MQTT

	urlBroker := mqtt.Protocol + "://" + mqtt.Url + ":" + mqtt.Port
	/*msg := mqttConfig.MessageSensorPublisher{
		sensor.Id,
		"temperature",
		sensor.AirportCode,
		"heure",
		12.1,
	}*/
	sensorId, err := strconv.Atoi(sensor.Id)
	QOSLevel, err := strconv.Atoi(sensor.QOSLevel)
	if err != nil {
		msg := mqttConfig.MessageSensorPublisher{
			SensorId:    sensorId,
			SensorType:  "temperature",
			AirportCode: sensor.AirportCode,
			Timestamp:   time.Now().Format("2017-09-07 17:06:06"),
			Value:       12.1,
		}
		client := mqttConfig.Connect(urlBroker, sensor.Id) //TODO
		client.Publish("aiport/capteur/temperature", byte(QOSLevel), true, msg)
	}

	//TODO finir le capteur température
}
