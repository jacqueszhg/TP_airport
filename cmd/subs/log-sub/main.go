package main

import (
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

	client := mqttConfig.Connect("tcp://localhost:1883", "log")
	client.Subscribe("airport/log", 1, myHandler)

	fmt.Printf("finish")
	wg.Wait()
}

func myHandler(client mqtt.Client, message mqtt.Message) {
	var res mqttConfig.MessageSensorPublisher
	err := json.Unmarshal(message.Payload(), &res)
	if err != nil {
		fmt.Println("Can't deserislize", message.Payload())
		fmt.Println(err)
	}

	pathFile := "log/%d.csv"
	date := res.Timestamp.Format("2006-01-02")
	pathFile = strings.ReplaceAll(pathFile, "%d", date)

	fmt.Println("\n")
	if !fileCSVExist(pathFile) {
		createCSV(pathFile)
	}
	openCSVAndWriteNewData(pathFile, res)
}

func fileCSVExist(filePath string) bool {
	// Check if the file exists
	if _, err := os.Stat(filePath); err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

func createCSV(filepath string) {
	// Create a new CSV file
	file, err := os.Create(filepath)
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	defer file.Close()

	// Change the permissions of the file "data.csv" to 644 (rw-r--r--)
	err = os.Chmod(filepath, 0755)
	if err != nil {
		// Handle error
		fmt.Print("permission : ")
		fmt.Println(err)
		return
	}

	// Create a new CSV writer
	writer := csv.NewWriter(file)

	// Write the header row
	header := []string{"Time", "Temperature", "Pressure", "Wind"}
	err = writer.Write(header)
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	// Flush the writer to ensure all data is written to the file
	writer.Flush()
}

func openCSVAndWriteNewData(filepath string, data mqttConfig.MessageSensorPublisher) {
	file, err := os.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			fmt.Println(err)

		} else {
			// Handle other error
			fmt.Println(err)
		}
		return
	}
	defer file.Close()

	// Create a new CSV writer
	// Create a new CSV reader
	writer := csv.NewWriter(file)

	newData := []string{
		data.Timestamp.Format("15:04:05"), fmt.Sprintf("%f", data.Value),
	}
	writer.Write(newData)

	// Flush the writer to ensure all data is written to the file
	writer.Flush()

	// Check for any errors
	err = writer.Error()
	if err != nil {
		// Handle erro
		fmt.Println(err)
		return
	}

	fmt.Println("Data written successfully")
}
