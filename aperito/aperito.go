package main

import (
	"gobot.io/x/gobot/platforms/mqtt"

	"log"

	"os"

	"time"
)

func main() {
	log.Println("Starting 'aperito'")
	if err := LoadConfig(); err != nil {
		log.Fatal("Failed loading configuration file: ", err)
	}
	log.Println("'aperito' config read")

	sender := Config.Mqtt.DefaultSender
	if len(os.Args) > 1 {
		// get the sender
		sender = os.Args[1]
	}

	//log.Println(Config.Mqtt.ConnectionString, Config.Mqtt.ClientName, Config.Mqtt.Topic)

	mqttAdaptor := mqtt.NewAdaptor(Config.Mqtt.ConnectionString, Config.Mqtt.ClientName)
	mqttAdaptor.SetAutoReconnect(true)

	if err := mqttAdaptor.Connect(); err != nil {
		log.Fatal("mqtt adaptor failed to connect", err)
	}

	log.Println("Publish result:", mqttAdaptor.Publish(Config.Mqtt.Topic, []byte(sender)))

	log.Println("sleeping...")
	time.Sleep(1000 * time.Millisecond)

	log.Println("Disconnecting")
	mqttAdaptor.Disconnect()

	log.Println("'aperito' says: Bye")
}
