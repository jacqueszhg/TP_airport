package pressure

import (
	mqttConfig "Airport/internal/pkg/mqtt"
	"fmt"
)

func main() {
	client := mqttConfig.Connect("tcp://localhost:1883", "aze")
	token := client.Publish("aiport/capteur/vent", 2, true, "je suis le vent pfou !")
	token.Wait()
	fmt.Printf("ok")
}
