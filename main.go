package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"simulator.iot.integrator.6th/src/infrastructure"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error al cargar variables de entorno:", err)
	}
	ctx, _ := context.WithCancel(context.Background())
	infrastructure.InitDependencies(
		ctx,
		os.Getenv("AMQP_URL"),
		os.Getenv("AMQP_EXCHANGE"),
	)
	sensorManager := infrastructure.GetSensorManager()

	sensorManager.InitSensorReadings(1)
	for {}
}