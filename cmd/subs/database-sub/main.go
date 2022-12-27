package main

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"time"
)

/*
 * SUB 1
 */

func main() {
	token := "IJERnyHz_xzjSMxSRh2lL1OO7IxXhBXj-0UFf3V2FOguLu-lINu_st8o4swU_005YJL8vD7oNAq24F8QWnZm3Q=="
	url := "https://europe-west1-1.gcp.cloud2.influxdata.com"
	client := influxdb2.NewClient(url, token)

	org := "yt.ryfax@gmail.com"
	bucket := "Tests"
	writeAPI := client.WriteAPIBlocking(org, bucket)

	for value := 0; value < 5; value++ {
		tags := map[string]string{
			"city": "NTE",
		}
		
		fields := map[string]interface{}{
			"value": value,
		}
		point := write.NewPoint("wind", tags, fields, time.Now())

		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		}
	}
}
