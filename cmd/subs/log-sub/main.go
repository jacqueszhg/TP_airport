package main

import (
	"Airport/internal/pkg/config"
	mqttConfig "Airport/internal/pkg/mqtt"
	"encoding/csv"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"strings"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	configSub := config.GetSubonfig("./logConfig.yml")

	// Connect to mqtt
	client := mqttConfig.Connect(configSub.MQTT.Protocol+"://"+configSub.MQTT.Url+":"+configSub.MQTT.Port, configSub.MQTT.Id)
	client.Subscribe("airport/log", 1, handler)
	fmt.Println("finish")
	wg.Wait()
}

func handler(client mqtt.Client, message mqtt.Message) {
	// Deserialize the object message
	var res mqttConfig.MessageSensorPublisher
	err := json.Unmarshal(message.Payload(), &res)
	if err != nil {
		fmt.Println("Can't deserislize", message.Payload())
		fmt.Println(err)
	}

	// Create the path for save the CSV
	pathFile := "../%airport/log/%sensorType/%airport-%date-%sensorType.csv"
	date := res.Timestamp.Format("2006-01-02")
	pathFile = strings.ReplaceAll(pathFile, "%sensorType", res.SensorType)
	pathFile = strings.ReplaceAll(pathFile, "%airport", res.AirportCode)
	pathFile = strings.ReplaceAll(pathFile, "%date", date)

	// Check if the file exist
	if !fileCSVExist(pathFile) {
		createCSV(pathFile, res.SensorType, res.UnitOfMeasure)
	}

	// Save data
	writeDataInCell(
		res,
		pathFile,
	)
}

func fileCSVExist(filePath string) bool {
	// Check if the file exists
	if _, err := os.Stat(filePath); err != nil {
		return false
	} else {
		return true
	}
}

func createCSV(filepath string, sensorType string, UnitOfMeasure string) {
	// Create a new CSV file
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	header := []string{
		"Time",
		sensorType + " (" + UnitOfMeasure + ")",
	}
	err = writer.Write(header)
	if err != nil {
		fmt.Println(err)
	}
	// Flush the writer to ensure all data is written to the file
	writer.Flush()

	// Check for any errors
	err = writer.Error()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("New file log is created at : ")
	fmt.Println(filepath)
}

func writeDataInCell(data mqttConfig.MessageSensorPublisher, filepath string) {
	// Open the file for reading
	file, err := os.Open(filepath)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read the file contents into a slice of slices of strings
	records, err := reader.ReadAll()
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}

	// Add new data in the array
	newRecord := []string{
		data.Timestamp.Format("15:04:05"),
		fmt.Sprintf("%f", data.Value),
	}
	records = append(records, newRecord)

	// Open the file for writing
	file, err = os.Create(filepath)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)

	// Write the modified data back to the file
	for _, record := range records {
		err = writer.Write(record)
		if err != nil {
			// Handle error
			return
		}
	}

	// Flush the writer to ensure all data is written to the file
	writer.Flush()

	// Check for any errors
	err = writer.Error()
	if err != nil {
		// Handle error
		return
	}

	fmt.Print("Data written successfully for : ")
	fmt.Print(data.SensorType)
	fmt.Print(" at ")
	fmt.Println(data.Timestamp)
}
