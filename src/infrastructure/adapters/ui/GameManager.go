package ui

import (
	"context"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameManager struct {
	resolution_w  int
	resolution_h  int
	value_ph      float32
	value_ntu     float32
	value_temp    float32
	value_tds     float32
	cancelContext context.CancelFunc
}

func NewGame() *GameManager {
	return &GameManager{
		value_ph:   0,
		value_ntu:  0,
		value_temp: 0,
		value_tds:  0,
	}
}

func (g *GameManager) Update() error {
	return nil
}

func (g *GameManager) Layout(outsideWidth, outsideHeight int) (int, int){
	return g.resolution_w, g.resolution_h
}

func (g *GameManager) Draw(screen *ebiten.Image){}