package main

import (
	mqttConfig "Airport/internal/pkg/mqtt"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	client := mqttConfig.Connect("tcp://localhost:1883", "sub")
	client.Subscribe("airport/temperature", 1, myHandler)

	fmt.Printf("finish")
	wg.Wait()
}

func myHandler(client mqtt.Client, message mqtt.Message) {
	// TODO : register in database
	var r mqttConfig.MessageSensorPublisher
	err := json.Unmarshal(message.Payload(), &r)
	if err != nil {
		fmt.Println("Can't deserislize", message.Payload())
	}

	fmt.Println(r)
	fmt.Println("\n")

}
