package ui


import (
	"github.com/hajimehoshi/ebiten/v2"
)

//Speed es cuantos ticks o frame del game 
type SpriteAnimation struct {
	Frames []*ebiten.Image
	Speed  int
	
	animFrame int
	animTick  int
}

func NewAnimation(frames []*ebiten.Image, speed int) *SpriteAnimation {
	return &SpriteAnimation{
		Frames: frames,
		Speed:  speed,
	}
}

//10/60 standar para cambiar el frame mío papu
func (a *SpriteAnimation) Update() {
	a.animTick++
	if a.animTick >= a.Speed {
		a.animFrame++
		if a.animFrame >= len(a.Frames) {
			a.animFrame = 0
		}
		a.animTick = 0
	}
}

func (a *SpriteAnimation) Draw(screen *ebiten.Image, x, y float64) {
	if a == nil || len(a.Frames) == 0 {
		return
	}
	
	currentFrame := a.Frames[a.animFrame]
	op := &ebiten.DrawImageOptions{}
	//lo va a mover a las pos que necesito 
	op.GeoM.Translate(x, y)
	screen.DrawImage(currentFrame, op)
}