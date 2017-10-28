#!/bin/sh
# GOARM=7 is for my Raspberry PI 3
GOARM=7 GOARCH=arm GOOS=linux go build -o godot
