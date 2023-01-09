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
		Id              string `yaml:"id"`
		Airport         string `yaml:"airport"`
		Frequency       string `yaml:"frequency"`
		QOSLevel        string `yaml:"QOSLevel"`
		AltitudeAirport string `yaml:"altitudeAirport"`
	}
}

type PubConfig struct {
	MQTT struct {
		Protocol string `yaml:"protocol"`
		Url      string `yaml:"url"`
		Port     string `yaml:"port"`
		Id       string `yaml:"id"`
	}
}

func GetSensorConfig(pathConfig string) SensorConfig {
	var config SensorConfig

	file, err := os.ReadFile(pathConfig)
	if err != nil {
		log.Fatal(err)
	}

	err2 := yaml.Unmarshal(file, &config)

	if err2 != nil {
		log.Fatal(err2)
	}
	return config
}

func GetSubonfig(pathConfig string) PubConfig {
	var config PubConfig

	file, err := os.ReadFile(pathConfig)
	if err != nil {
		log.Fatal(err)
	}

	err2 := yaml.Unmarshal(file, &config)

	if err2 != nil {
		log.Fatal(err2)
	}
	return config
}
