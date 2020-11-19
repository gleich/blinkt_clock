package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Matt-Gleich/logoru"
)

var (
	timeLevel  string
	lightLevel int
)

func main() {
	parseArgs()
	getLevel()
	fmt.Println(lightLevel)
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
}
