#!/bin/sh
# go get -d -u github.com/m4ng0/go-rfid-rc522/rfid/...
# GOARM=7 is for my Raspberry PI 3
CC=arm-linux-gnueabihf-gcc CGO_ENABLED=1 GOARM=7 GOARCH=arm GOOS=linux go build
