package database_sub

import (
	mqttConfig "Airport/internal/pkg/mqtt"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	client := mqttConfig.Connect("tcp://localhost:1883", "azazeazee")
	client.Subscribe("aiport/capteur/vent", 2, myHandler)

	fmt.Printf("finish")
	wg.Wait()
}

func myHandler(client mqtt.Client, message mqtt.Message) {
	// TODO :Register sensors data in influxDB
	fmt.Println(string(message.Payload()))
}
