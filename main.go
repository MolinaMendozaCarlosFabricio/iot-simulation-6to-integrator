package main

import (

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
	ctx, _ := context.WithCancel(context.Background())
	infrastructure.InitDependencies(
		ctx,
		os.Getenv("AMQP_URL"),
		os.Getenv("AMQP_EXCHANGE"),
	)
	sensorManager := infrastructure.GetSensorManager()

	game := ui.NewGame()
	log.Println("icniando los snesoressss...")
	go sensorManager.InitSensorReadings(1, game)	

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Simulador IoT - AquaFlow)")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}