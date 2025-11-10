package infrastructure

import (
	"context"
	"fmt"

	"simulator.iot.integrator.6th/src/infrastructure/adapters/queue"
	"simulator.iot.integrator.6th/src/infrastructure/adapters/sensors"
)

var sensorManager sensors.SimulatedSensorsManager
var rabbitMQManager queue.RabbitMQManager

func InitDependencies(
	ctx context.Context,
	url,
	exchange string,
){
	rabbitMQManager, err := queue.NewRabbitMQManager(url, exchange)
	if err != nil {
		fmt.Println("Error al inicializar rabbitMQ:", err)
	}
	sensorManager = *sensors.NewSensorsManager(ctx, *rabbitMQManager)
}

func GetSensorManager()*sensors.SimulatedSensorsManager{
	return &sensorManager
}

func GetRabbitMQManager()*queue.RabbitMQManager{
	return &rabbitMQManager
}