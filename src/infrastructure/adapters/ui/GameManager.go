package ui

import (
	"context"
	"log"
	"sync"

	"fmt"
	"image/color"

	"simulator.iot.integrator.6th/src/application"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameManager struct {
	resolution_w  int
	resolution_h  int
	value_ph      float32
	value_ntu     float32
	value_temp    float32
	value_tds     float32
	cancelContext context.CancelFunc
	finish		  bool
	exitStatus	  error

	stateMutex sync.RWMutex

	//inyectada
	
	latestData application.SensorDisplayDTO
	
	//sprites, fondos, img, etc.
	iotImg *ebiten.Image
	waterBucketImg *ebiten.Image
	sensorSprites map[string]*ebiten.Image
	
	
	

	//animacion de sprites
	bdSpritesAnim  *SpriteAnimation
	awsSpritesAnim *SpriteAnimation

	iotSpritesAnim *SpriteAnimation
	tempSpritesAnim *SpriteAnimation
	phSpritesAnim *SpriteAnimation
	tdsSpritesAnim *SpriteAnimation
	ntuSpritesAnim *SpriteAnimation

	//animacion de flechas

	tempAnim *SpriteAnimation
	phAnim   *SpriteAnimation
	tdsAnim  *SpriteAnimation
	ntuAnim  *SpriteAnimation
	awsAnim *SpriteAnimation
}

func NewGame(cancelCtx context.CancelFunc, exitStatus error) *GameManager {	
	
	//imagnes y sprites

	//bucket
	bucketImg, _, err := ebitenutil.NewImageFromFile("assets/iot/bucket.png")
	if err != nil {
		log.Fatalf("Error al cargar wl water: %v", err)
	}


	//ANIMCACIONES DE SPRITES

	//bd
	bdSprites := []* ebiten.Image{
		MustNewImageFromFile("assets/db/db_1-new.png"),
		MustNewImageFromFile("assets/db/db_2-new.png"),
	}

		//aws
	awsprites := []* ebiten.Image{
		MustNewImageFromFile("assets/aws/aws_1-new.png"),
		MustNewImageFromFile("assets/aws/aws_2-new.png"),
		MustNewImageFromFile("assets/aws/aws_3-new.png"),
	}

	//iot
	iotSprites := []* ebiten.Image{
		MustNewImageFromFile("assets/iot/iot_1-new.png"),
		MustNewImageFromFile("assets/iot/iot_2-new.png"),
		MustNewImageFromFile("assets/iot/iot_3-new.png"),
	}

		//temp
	tempSprites := []* ebiten.Image{
		MustNewImageFromFile("assets/temp_sensor/temp_sensor_1-new.png"),
		MustNewImageFromFile("assets/temp_sensor/temp_sensor_2-new.png"),
		MustNewImageFromFile("assets/temp_sensor/temp_sensor_3-new.png"),
	}

		//ph
	phSprites := []* ebiten.Image{
		MustNewImageFromFile("assets/ph_sensor/ph_sensor_1-new.png"),
		MustNewImageFromFile("assets/ph_sensor/ph_sensor_2-new.png"),
		MustNewImageFromFile("assets/ph_sensor/ph_sensor_3-new.png"),
	}


		//tds
	tdsSprites := []* ebiten.Image{
		MustNewImageFromFile("assets/tds_sensor/tds_sensor_1-new.png"),
		MustNewImageFromFile("assets/tds_sensor/tds_sensor_2-new.png"),
		MustNewImageFromFile("assets/tds_sensor/tds_sensor_3-new.png"),
	}


		//ntu
	ntuSprites := []* ebiten.Image{
		MustNewImageFromFile("assets/ntu_sensor/ntu_sensor_1-new.png"),
		MustNewImageFromFile("assets/ntu_sensor/ntu_sensor_2-new.png"),
		MustNewImageFromFile("assets/ntu_sensor/ntu_sensor_3-new.png"),
	}



	//ANIMACIONES DE FLECHAS
	

		// TEMP
	tempFrames := []*ebiten.Image{
		MustNewImageFromFile("assets/temp_sensor/arrowA1.png"),
		MustNewImageFromFile("assets/temp_sensor/arrowA2.png"),
		MustNewImageFromFile("assets/temp_sensor/arrowA3.png"),
	}

	
	phFrames := []*ebiten.Image{
		MustNewImageFromFile("assets/ph_sensor/arrowA1.png"),
		MustNewImageFromFile("assets/ph_sensor/arrowA2.png"),
		MustNewImageFromFile("assets/ph_sensor/arrowA3.png"),
		MustNewImageFromFile("assets/ph_sensor/arrowA4.png"),
		MustNewImageFromFile("assets/ph_sensor/arrowA5.png"),
		MustNewImageFromFile("assets/ph_sensor/arrowA6.png"),
	}
	
	// TDS
	tdsFrames := []*ebiten.Image{
		MustNewImageFromFile("assets/tds_sensor/arrowA1.png"),
		MustNewImageFromFile("assets/tds_sensor/arrowA2.png"),
		MustNewImageFromFile("assets/tds_sensor/arrowA3.png"),

		
	}

	
	// NTU
	ntuFrames := []*ebiten.Image{
		MustNewImageFromFile("assets/ntu_sensor/arrowA1.png"),
		MustNewImageFromFile("assets/ntu_sensor/arrowA2.png"),
		MustNewImageFromFile("assets/ntu_sensor/arrowA3.png"),
		MustNewImageFromFile("assets/ntu_sensor/arrowA4.png"),
		MustNewImageFromFile("assets/ntu_sensor/arrowA5.png"),
		MustNewImageFromFile("assets/ntu_sensor/arrowA6.png"),
	}
	
	// AWS
	awsFrames := []*ebiten.Image{
		MustNewImageFromFile("assets/aws/arrowA1.png"),
		MustNewImageFromFile("assets/aws/arrowA2.png"),
		MustNewImageFromFile("assets/aws/arrowA3.png"),
	}
	
	
	return &GameManager{
		value_ph:   0,
		value_ntu:  0,
		value_temp: 0,
		value_tds:  0,

		 
		// iotImg: deviceImg,
		// sensorSprites: sprites,
		waterBucketImg: bucketImg,
		
		//animacion es flechas
		tempAnim: NewAnimation(tempFrames, 10),
		phAnim: NewAnimation(phFrames, 10),
		tdsAnim: NewAnimation(tdsFrames, 10),
		ntuAnim: NewAnimation(ntuFrames, 10),
		awsAnim: NewAnimation(awsFrames, 10),

		//animaciones sprites
		bdSpritesAnim: NewAnimation(bdSprites, 20),
		awsSpritesAnim: NewAnimation(awsprites,20),

		iotSpritesAnim: NewAnimation(iotSprites,20),
		tempSpritesAnim: NewAnimation(tempSprites,20),
		phSpritesAnim: NewAnimation(phSprites,20),
		tdsSpritesAnim: NewAnimation(tdsSprites,20),
		ntuSpritesAnim: NewAnimation(ntuSprites,20),




		resolution_w:  1280, 
        resolution_h:  720,
		cancelContext: cancelCtx,
		finish: false,
		exitStatus: exitStatus,
	}
}

func (g *GameManager) Update() error {

	g.tempAnim.Update()
    g.phAnim.Update()
	g.tdsAnim.Update()
	g.ntuAnim.Update()
	g.awsAnim.Update()

	//animaciones de sprites
	g.bdSpritesAnim.Update()
	g.awsSpritesAnim.Update()

	g.iotSpritesAnim.Update()
	g.tempSpritesAnim.Update()
	g.phSpritesAnim.Update()
	g.tdsSpritesAnim.Update()
	g.ntuSpritesAnim.Update()




	// para poder salir 
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || g.finish {
		println("Cancelando contexto y cerrando juego...")
		// acá cierra los goroutines 
		g.cancelContext()
		return g.exitStatus
	}
	return nil
}

func (g *GameManager) UpdateState(dto application.SensorDisplayDTO) {
	// para poder escrinir
	g.stateMutex.Lock() 
	defer g.stateMutex.Unlock()

	//save el dto
	g.latestData = dto
}


func (g *GameManager) Layout(outsideWidth, outsideHeight int) (int, int){
	return g.resolution_w, g.resolution_h
}

func (g *GameManager) Draw(screen *ebiten.Image){
	
	
	screen.Fill(color.White)

	//para ller 
	g.stateMutex.RLock() 
	currentData := g.latestData 
	g.stateMutex.RUnlock() 


	//iot en medio
	if g.iotImg != nil {
		opDevice := &ebiten.DrawImageOptions{}
		opDevice.GeoM.Translate(float64(512), float64(52))
		screen.DrawImage(g.iotImg, opDevice)
	}

	//agua
	if g.waterBucketImg != nil {
		opBucket := &ebiten.DrawImageOptions{}
		opBucket.GeoM.Translate(float64(480), float64(400))
		screen.DrawImage(g.waterBucketImg, opBucket)
	}


	safeText := "NO REUTILIZABLE"
    if currentData.IsWaterSafe {
        safeText = "REUTILIZABLE"
    }

	debugText := fmt.Sprintf(
		"ESTADO: %s\n\nPH: %.2f (%s)\nTEMP: %.2f (%s)\nTDS: %.2f (%s)\nTURB: %.2f (%s)",
		safeText,
		currentData.PHValue, currentData.PHState,
		currentData.TempValue, currentData.TempState,
		currentData.TDSValue, currentData.TDSState,
		currentData.TurbValue, currentData.TurbState,
	)
	ebitenutil.DebugPrint(screen, debugText)


	//animaciones

	//flechas
	g.tempAnim.Draw(screen, float64(700), float64(260))
	g.phAnim.Draw(screen, float64(619), float64(250)) 
	g.tdsAnim.Draw(screen, float64(438), float64(240))
	g.ntuAnim.Draw(screen, float64(393), float64(234))
	g.awsAnim.Draw(screen, float64(752), float64(80))

	//sprites
	g.bdSpritesAnim.Draw(screen, float64(201), float64(-4))
	g.awsSpritesAnim.Draw(screen, float64(888), float64(-4))

	g.iotSpritesAnim.Draw(screen, float64(512), float64(52))

	g.tempSpritesAnim.Draw(screen, float64(777), float64(188))

	g.phSpritesAnim.Draw(screen, float64(768), float64(428))

	g.tdsSpritesAnim.Draw(screen, float64(272), float64(188))

	g.ntuSpritesAnim.Draw(screen, float64(272), float64(428))


}


func MustNewImageFromFile(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatalf("Error al cargar imagen (Must): %s - %v", path, err)
	}
	return img
}