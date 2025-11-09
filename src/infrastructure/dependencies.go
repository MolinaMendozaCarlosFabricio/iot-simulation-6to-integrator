package infrastructure

import (
	"context"

	"simulator.iot.integrator.6th/src/infrastructure/adapters/sensors"
)

var sensorManager sensors.SimulatedSensorsManager

func InitDependencies(ctx context.Context){
	sensorManager = *sensors.NewSensorsManager(ctx)
}

func GetSensorManager()*sensors.SimulatedSensorsManager{
	return &sensorManager
}