package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/audio"
	"gobot.io/x/gobot/platforms/raspi"

	"log"

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

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			log.Println("ring button pressed")
			ring.Play()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			log.Println("ring button released")
		})
	}

	robot := gobot.NewRobot("Godot Robot",
		[]gobot.Connection{r, a},
		[]gobot.Device{button, ring},
		work,
	)

	robot.Start()
}
