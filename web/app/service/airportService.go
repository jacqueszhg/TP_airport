package service

import (
	"Airport/web/app/model"
	"Airport/web/app/repository"
	"time"
)

func GetMeasuresByAirportAndType(airportCode string, measurement string, startDate time.Time, endDate time.Time) []model.Measure {

	//convert the date to match in query
	start := startDate.Format("2006-01-02T15:04:05Z")
	stop := endDate.Format("2006-01-02T15:04:05Z")

	// we get measures from repository
	result := repository.GetMeasuresByAirportAndType(airportCode, measurement, start, stop)

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

	// get all measures for a specific time range from repository
	result := repository.GetAllMeasuresByDate(airportCode, date)

	// to calculate average of each sensor type
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

	//calculate average of each sensor type
	temperatureAverage := 0.0
	if temperatureCompt != 0.0 {
		temperatureAverage = temperatureSomme / temperatureCompt
	}
	pressureAverage := 0.0
	if pressureCompt != 0.0 {
		pressureAverage = pressureSomme / pressureCompt
	}
	windAverage := 0.0
	if windCompt != 0.0 {
		windAverage = windSomme / windCompt
	}

	return temperatureAverage, pressureAverage, windAverage
}

func GetAllAirports() []string {

	// get all airport fields
	result := repository.GetAllAirports()

	var liste []string

	for result.Next() {
		record := result.Record().ValueByKey("_value").(string)
		liste = append(liste, record)
	}

	return liste

}
