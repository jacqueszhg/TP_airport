package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type SensorConfig struct {
	MQTT struct {
		Protocol string `yaml:"protocol"`
		Url      string `yaml:"url"`
		Port     string `yaml:"port"`
	}

	Sensor struct {
		Id          string `yaml:"id"`
		AirportCode string `yaml:"airportCode"`
		Frequency   string `yaml:"frequency"`
		QOSLevel    string `yaml:"QOSLevel"`
	}
}

func GetSensorConfig() SensorConfig {
	var config SensorConfig

	file, err := os.ReadFile("configs/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	err2 := yaml.Unmarshal(file, &config)

	if err2 != nil {
		log.Fatal(err2)
	}

	return config
}
