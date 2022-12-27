package model

type Measure struct {
	Id          string  `json:"_id,omitempty" bson:"_id,omitempty"`
	SensorId    string  `json:"sensorid"`
	AirportCode string  `json:"airportid"`
	Timestamp   string  `json:"date"`
	Value       float64 `json:"value"`
	SensorType  string  `json:"sensortype"`
}

type Average struct {
	SensorType string  `json:"sensortype"`
	Average    float64 `json:"average"`
}

type Airport struct {
	Id   string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
	City string `json:"city"`
}
