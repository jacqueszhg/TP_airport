package controller

import (
	"Airport/web/app/helper"
	"Airport/web/app/model"
	"Airport/web/app/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @BasePath /api/v1

// GetMeasures godoc
// @Summary Return a list of value for one type (temperature, wind, pressure) between two time
// @Description Get measurements of a certain type (temperature, wind, pressure) that are between two time (date + time)
// @Schemes
// @Tags airport
// @Accept json
// @Produce json
// @Param   airportCode     path    string  true         	"airport code IATA"
// @Param 	type 			query 	string 	true 			"sensor type (temperature, wind, pressure)"
// @Param 	startDate 		query 	string 	true 			"start date (example : 2021-04-04T22:08:41Z)"
// @Param 	endDate 		query 	string 	false 			"end date (example : 2021-04-04T22:08:41Z)"
// @Success 200 {array} model.Measure
// @Failure 400 {object} helper.ErrorResponse
// @Router /airport/{airportCode}/measure [get]
func GetMeasures(g *gin.Context) {
	airportCode, _ := g.Params.Get("airportCode")

	sensorType, isPresent := g.GetQuery("type")
	if !isPresent {
		helper.GetError(
			errors.New("Missing type in query"),
			g.Writer,
			400,
		)
		return
	}

	startDate, isPresent := g.GetQuery("startDate")
	if !isPresent {
		helper.GetError(
			errors.New("Missing startDate in query"),
			g.Writer,
			400,
		)
		return
	}
	endDate, isPresent := g.GetQuery("endDate")

	startDateConvert, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		fmt.Println(err)
		helper.GetError(
			errors.New("Invalid dateStart : correct format : 2021-04-04T22:08:41Z"),
			g.Writer,
			400,
		)
		return
	}
	endDateConvert, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		helper.GetError(
			errors.New("Invalid dateEnd : correct format : 2021-04-04T22:08:41Z"),
			g.Writer,
			400,
		)
		return
	}

	res := service.GetMeasuresByAirportAndType(airportCode, sensorType, startDateConvert, endDateConvert)
	g.JSON(http.StatusOK, res)
}

// GetAverages godoc
// @Summary Return three averages (temperature, pressure, wind) for a specific date
// @Description Get averages of measures (temperature, pressure, wind) for a specific date
// @Schemes
// @Tags airport
// @Accept json
// @Produce json
// @Param   airportCode     path    string  true         	"airport code IATA"
// @Param 	date	query 	string 	true  "start date (example : 2021-04-04)"
// @Success 200 {array} model.Average
// @Failure 400 {object} helper.ErrorResponse
// @Router /airport/{airportCode}/averages [get]
func GetAverages(g *gin.Context) {
	airportCode, _ := g.Params.Get("airportCode")

	date, isPresent := g.GetQuery("date")
	if !isPresent {
		helper.GetError(
			errors.New("Missing startDate in query"),
			g.Writer,
			400,
		)
		return
	}

	t, err := time.Parse("2006-01-02", date)
	if err == nil {
		temperatureAverage, pressureAverage, windAverage := service.GetAveragesByDate(airportCode, t)

		g.JSON(http.StatusOK, []model.Average{
			model.Average{
				SensorType: "temperature",
				Average:    temperatureAverage,
				Unit:       "°C",
			},
			model.Average{
				SensorType: "pressure",
				Average:    pressureAverage,
				Unit:       "hPa",
			},
			model.Average{
				SensorType: "wind",
				Average:    windAverage,
				Unit:       "m/s",
			},
		})
	} else {
		helper.GetError(
			errors.New("Invalid date : correct format : 2021-04-04"),
			g.Writer,
			400,
		)
		return
	}

}

// GetAllAirport godoc
// @Summary Return all the airport in DB
// @Description Return all the airport in DB
// @Schemes
// @Tags airport
// @Produce json
// @Success 200 {array} string
// @Failure 400 {object} helper.ErrorResponse
// @Router /airports [get]
func GetAllAirport(g *gin.Context) {

	//liste des codes IATA présent en base
	var liste = service.GetAllAirports()

	g.JSON(http.StatusOK, liste)

	return
}
