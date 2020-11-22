package main

import (
	"os"
	"time"

	"github.com/Matt-Gleich/logoru"
	blinkt "github.com/alexellis/blinkt_go"
)

var (
	timeLevel  string
	lightLevel int
	lights     = blinkt.NewBlinkt(1.0)
)

func main() {
	lights.SetClearOnExit(true)
	lights.Setup()
	parseArgs()
	for {
		getLevel()
		setLevel()
		time.Sleep(time.Millisecond * 1)
	}
}

func parseArgs() {
	timeLevel = os.Args[0]
	logoru.Info("Got time level")
}

// Get the light level
func getLevel() {
	var (
		now        = time.Now()
		percentage float32
	)
	switch timeLevel {
	case "hours":
		percentage = float32(now.Hour()) / 24
	case "minutes":
		percentage = float32(now.Minute()) / 60
	default:
		percentage = float32(now.Second()) / 60
	}
	lightLevel = int(percentage * 10)
	logoru.Info("Got time level:", percentage)
}

// Set the light level
func setLevel() {
	lights.Clear()
	for pixel := 0; pixel < lightLevel-1; pixel++ {
		lights.SetPixel(pixel, 0, 255, 0)
	}
	lights.Show()
	logoru.Info("Updated lights")
}
