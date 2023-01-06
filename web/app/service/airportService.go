package service

import (
	"Airport/web/app/model"
	"context"
	"fmt"
	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"time"
)

func getDbClient() influxdb.Client {
	token := "IJERnyHz_xzjSMxSRh2lL1OO7IxXhBXj-0UFf3V2FOguLu-lINu_st8o4swU_005YJL8vD7oNAq24F8QWnZm3Q=="
	url := "https://europe-west1-1.gcp.cloud2.influxdata.com"
	return influxdb.NewClient(url, token)
}

func getQueryAPI(client influxdb.Client) api.QueryAPI {
	org := "airport"
	return client.QueryAPI(org)
}

func GetMeasuresByAirportAndType(airportCode string, measurement string, startDate time.Time, endDate time.Time) []model.Measure {
	client := getDbClient()
	queryAPI := getQueryAPI(client)

	bucket := "Sensors"
	start := startDate.Format("2006-01-02T15:04:05Z")
	stop := endDate.Format("2006-01-02T15:04:05Z")

	query := fmt.Sprintf(`from(bucket: "%v") |> range(start: %v, stop: %v) |> filter(fn: (r) => r["airport"] == "%v") |> filter(fn: (r) => r["_measurement"] == "%v")`, bucket, start, stop, airportCode, measurement)
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	var res []model.Measure
	for result.Next() {
		record := result.Record()
		time := record.Time().Format("2006-01-02T15:04:05Z")

		measure := model.Measure{
			SensorId:    record.ValueByKey("id").(string),
			AirportCode: airportCode,
			Timestamp:   time,
			Value:       record.Value().(float64),
			SensorType:  record.Measurement(),
		}

		res = append(res, measure)
	}

	return res
}

func GetAveragesByDate(airportCode string, date time.Time) (float64, float64, float64) {
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

	// temperature
	temperatureCompt, temperatureSomme := 0.0, 0.0
	pressureCompt, pressureSomme := 0.0, 0.0
	windCompt, windSomme := 0.0, 0.0

	for result.Next() {
		record := result.Record()
		if record.Measurement() == "temperature" {
			temperatureSomme += record.Value().(float64)
			temperatureCompt++
		} else if record.Measurement() == "pressure" {
			pressureSomme += record.Value().(float64)
			pressureCompt++
		} else if record.Measurement() == "wind" {
			windSomme += record.Value().(float64)
			windCompt++
		}
	}

	temperatureAverage := temperatureSomme / temperatureCompt
	pressureAverage := pressureSomme / pressureCompt
	windAverage := windSomme / windCompt

	client.Close()

	return temperatureAverage, pressureAverage, windAverage
}

func GetAllAirports() []string {

	client := getDbClient()
	queryAPI := getQueryAPI(client)
	query := fmt.Sprintf(`import "influxdata/influxdb/v1" v1.tagValues(bucket: "Sensors", tag:  "airport" )`)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	var liste []string

	for result.Next() {
		record := result.Record().ValueByKey("_value").(string)
		liste = append(liste, record)
	}
	client.Close()
	return liste

}
