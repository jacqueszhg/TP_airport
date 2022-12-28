package main

import (
	mqttConfig "Airport/internal/pkg/mqtt"
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Initialize Influx DB
	db := createDb()
	writeApi := createWriteAPI(db)

	// Connect to mqtt
	client := mqttConfig.Connect("tcp://localhost:1883", "sub")

	// Subscribe to all sensors
	client.Subscribe("airport/temperature", 1, func(client mqtt.Client, message mqtt.Message) {
		onDataReceived(message, writeApi)
	})

	fmt.Printf("finish")
	wg.Wait()
}

/*
 * Creating the influx database with the corresponding token and url
 */
func createDb() influxdb.Client {
	token := "IJERnyHz_xzjSMxSRh2lL1OO7IxXhBXj-0UFf3V2FOguLu-lINu_st8o4swU_005YJL8vD7oNAq24F8QWnZm3Q=="
	url := "https://europe-west1-1.gcp.cloud2.influxdata.com"
	return influxdb.NewClient(url, token)
}

/*
 * Creating a WriteAPIBlocking with our database and organization/bucket
 */
func createWriteAPI(db influxdb.Client) api.WriteAPIBlocking {
	org := "yt.ryfax@gmail.com"
	bucket := "Sensors"
	return db.WriteAPIBlocking(org, bucket)
}

/*
 * On any data received (from every sensors), add a point into the database
 */
func onDataReceived(message mqtt.Message, api api.WriteAPIBlocking) {
	var r mqttConfig.MessageSensorPublisher
	err := json.Unmarshal(message.Payload(), &r)
	if err != nil {
		fmt.Println("Can't deserislize", message.Payload())
	}

	tags := map[string]string{
		"airport": r.AirportCode,
		"id":      string(rune(r.SensorId)),
	}

	fields := map[string]interface{}{
		"value": r.Value,
	}

	point := write.NewPoint(r.SensorType, tags, fields, time.Now())

	if err := api.WritePoint(context.Background(), point); err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)
	fmt.Println("Sent!")
}
