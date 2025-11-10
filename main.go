package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"simulator.iot.integrator.6th/src/application"
	"simulator.iot.integrator.6th/src/infrastructure/adapters/ui"
)

func main() {
    // 1. Crea el "Mock" (el simulador falso)
    appService := application.NewMockSensorService()
    
    // 2. Pásaselo a tu UI
    // (Tu UI cree que es el servicio real, ¡esa es la magia!)
    game := ui.NewGame(appService)
    
	ebiten.SetWindowSize(1280, 720) 
	ebiten.SetWindowTitle("Simulador IoT (Usando Mock)")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}