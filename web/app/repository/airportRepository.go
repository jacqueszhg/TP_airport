package repository

import (
	"context"
	"fmt"
	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"time"
)

// we create a new influxDB client to connect with the database with token and url
func getDbClient() influxdb.Client {
	token := "IJERnyHz_xzjSMxSRh2lL1OO7IxXhBXj-0UFf3V2FOguLu-lINu_st8o4swU_005YJL8vD7oNAq24F8QWnZm3Q=="
	url := "https://europe-west1-1.gcp.cloud2.influxdata.com"
	return influxdb.NewClient(url, token)
}

// we create a new QueryAPI to can get data in with flux
func getQueryAPI(client influxdb.Client) api.QueryAPI {
	org := "airport"
	return client.QueryAPI(org)
}

// GetMeasuresByAirportAndType retrieve the data of a type of sensors between two distinct dates
func GetMeasuresByAirportAndType(airportCode string, measurement string, start string, stop string) *api.QueryTableResult {
	client := getDbClient()
	queryAPI := getQueryAPI(client)

	bucket := "Sensors"
	query := fmt.Sprintf(`from(bucket: "%v") |> range(start: %v, stop: %v) |> filter(fn: (r) => r["airport"] == "%v") |> filter(fn: (r) => r["_measurement"] == "%v")`, bucket, start, stop, airportCode, measurement)
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	client.Close()
	return result
}

// GetAllMeasuresByDate retrieve all values for a range
func GetAllMeasuresByDate(airportCode string, date time.Time) *api.QueryTableResult {
	client := getDbClient()
	queryAPI := getQueryAPI(client)

	bucket := "Sensors"

	start := date.Format("2006-01-02T15:04:05Z")
	stop := date.AddDate(0, 0, 1).Format("2006-01-02T15:04:05Z")

	query := fmt.Sprintf(`from(bucket: "%v") |> range(start: %v, stop: %v) |> filter(fn: (r) => r["airport"] == "%v")`, bucket, start, stop, airportCode)
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	client.Close()
	return result
}

// GetAllAirports get all airport fields
func GetAllAirports() *api.QueryTableResult {
	client := getDbClient()
	queryAPI := getQueryAPI(client)
	query := fmt.Sprintf(`import "influxdata/influxdb/v1" v1.tagValues(bucket: "Sensors", tag:  "airport" )`)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	client.Close()
	return result
}
