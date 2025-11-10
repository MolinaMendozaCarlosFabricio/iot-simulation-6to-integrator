package main

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"simulator.iot.integrator.6th/src/infrastructure/adapters/ui"

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
	ctx, cancelFunc := context.WithCancel(context.Background())
	infrastructure.InitDependencies(
		ctx,
		os.Getenv("AMQP_URL"),
		os.Getenv("AMQP_EXCHANGE"),
	)
	sensorManager := infrastructure.GetSensorManager()

	exitStatus := errors.New("exit game")
	game := ui.NewGame(cancelFunc, exitStatus)
	log.Println("icniando los sensores...")
	go sensorManager.InitSensorReadings(1, game)	

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Simulador IoT - AquaFlow")

	if err := ebiten.RunGame(game); err != nil && err != exitStatus {
		log.Fatal(err)
	}

}