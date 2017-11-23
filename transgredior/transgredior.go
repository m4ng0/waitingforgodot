package main

import (
	"gobot.io/x/gobot/platforms/audio"

	"log"

	rfid "github.com/m4ng0/go-rfid-rc522/rfid"
	rc522 "github.com/m4ng0/go-rfid-rc522/rfid/rc522"

	"os/exec"

	"time"
)

func main() {
	log.Println("Starting 'transgredior'")
	if err := LoadConfig(); err != nil {
		log.Fatal("Failed loading configuration file: ", err)
	}

	reader, err := rc522.NewRfidReader()
	if err != nil {
		log.Fatal(err)
		return
	}
	readerChan, err := rfid.NewReaderChan(reader, 100*time.Millisecond)
	if err != nil {
		log.Fatal(err)
		return
	}
	rfidChan := readerChan.GetChan()

	a := audio.NewAdaptor()
	grantedSound := audio.NewDriver(a, Config.Rfid.Sound.AccessGranted)
	deniedSound := audio.NewDriver(a, Config.Rfid.Sound.AccessDenied)
	a.Connect()
	grantedSound.Start()
	deniedSound.Start()

	verificationChannel := make(chan string, 1) // one slot buffer
	for {
		id := <-rfidChan
		log.Println("Read: ", id)
		select {
		case verificationChannel <- id: // put the value in the channel, if it isn't full
			go func() {
				defer func() {
					<-verificationChannel
				}()
				found := false
				for _, a := range Config.Rfid.Access {
					if id == a.Key {
						found = true
						break
					}
				}
				if found {
					log.Println("Access granted")
					grantedSound.Play()
					cmd := exec.Command(Config.Rfid.Command.AccessGranted, id)
					if err := cmd.Run(); err != nil {
						log.Println("Could not execute command!")
					}
				} else {
					log.Printf("Got id %s: someone is trying to do something nasty?", id)
					deniedSound.Play()
				}
				time.Sleep(1000 * time.Millisecond)
				log.Println("Ready for next")
			}()
		default:
			log.Println("Channel full, verification discarded")
		}
	}
}
