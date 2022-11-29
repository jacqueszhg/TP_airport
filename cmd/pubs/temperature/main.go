package main

import (
	"Airport/internal/pkg/config"
	mqttConfig "Airport/internal/pkg/mqtt"
)

func main() {

	sensor := config.GetSensorConfig().Sensor
	mqtt := config.GetSensorConfig().MQTT

	urlBroker := mqtt.Protocol + "://" + mqtt.Url + ":" + mqtt.Port
	msg := mqttConfig.MessageSensorPublisher{
		sensor.Id,
		"temperature",
		sensor.AirportCode,
		"heure",
		12.1,
	}

	client := mqttConfig.Connect(urlBroker, sensor.Id) //TODO
	token := client.Publish("aiport/capteur/temperature", 1, true, msg)
	token.Wait()

	//TODO finir le capteur température
}
