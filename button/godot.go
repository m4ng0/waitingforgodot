package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/audio"
	"gobot.io/x/gobot/platforms/raspi"

	"log"

	"os/exec"

	"time"
)

func main() {
	log.Println("Starting Godot")
	if err := LoadConfig(); err != nil {
		log.Fatal("Failed loading configuration file: ", err)
	}

	r := raspi.NewAdaptor()
	button := gpio.NewButtonDriver(r, Config.Button.Pin, time.Duration(Config.Button.PollInterval)*time.Millisecond)
	button.DefaultState = Config.Button.DefaultState

	a := audio.NewAdaptor()
	ring := audio.NewDriver(a, Config.Button.PressSoundFile)

	r.Connect()
	button.Start()
	a.Connect()
	ring.Start()

	buttonEvents := button.Subscribe()

	pressEventChannel := make(chan *gobot.Event, 1) // only one event buffer
	for {
		select {
		case event := <-buttonEvents:
			log.Println("Event: ", event.Name, event.Data)
			if event.Name == "push" {
				log.Println("ring button pressed")
				select {
				case pressEventChannel <- event: // put the event in the channel, if it isn't full
					go func() {
						defer func() {
							<-pressEventChannel
						}()
						ring.Play()
						cmd := exec.Command(Config.Button.Command)
						if err := cmd.Run(); err != nil {
							log.Println("Could not execute command!")
						}
					}()
				default:
					log.Println("Channel full, discarding press event")
				}
			} else if event.Name == "release" {
				log.Println("ring button released")
			}
		}
	}
}
