package main

import (

	"github.com/hajimehoshi/ebiten/v2"
	"simulator.iot.integrator.6th/src/application"
	"simulator.iot.integrator.6th/src/infrastructure/adapters/ui"

	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"simulator.iot.integrator.6th/src/infrastructure"
)



func main() {
    appService := application.NewMockSensorService()

    game := ui.NewGame(appService)

	ebiten.SetWindowSize(1280, 720) 
	ebiten.SetWindowTitle("Simulador IoT (Usando Mock)")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
	
	
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