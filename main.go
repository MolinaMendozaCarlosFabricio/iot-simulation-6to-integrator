package main

import (
	"context"

	"simulator.iot.integrator.6th/src/infrastructure"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	infrastructure.InitDependencies(ctx)
	sensorManager := infrastructure.GetSensorManager()

	sensorManager.InitSensorReadings(1)
	for {}
}