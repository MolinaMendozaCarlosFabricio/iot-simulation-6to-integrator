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

	//animacion de flechas

	tempAnim *SpriteAnimation
	// phAnim   *SpriteAnimation
	// tdsAnim  *SpriteAnimation
	// ntuAnim  *SpriteAnimation
}

func NewGame(cancelCtx context.CancelFunc, exitStatus error) *GameManager {	
	
	//imagnes y sprites

	//iot
	deviceImg, _, err := ebitenutil.NewImageFromFile("assets/iot/iot_1-new.png") 
    if err != nil {
		log.Fatalf("Error al cargar iot_device.png: %v", err)
	}

	bucketImg, _, err := ebitenutil.NewImageFromFile("assets/iot/bucket.png")
	if err != nil {
		log.Fatalf("Error al cargar wl water: %v", err)
	}
	
	sprites := make(map[string]*ebiten.Image)

	//ph
	sprites["PH_NORMAL"], _, _ = ebitenutil.NewImageFromFile("assets/ph_sensor/ph_sensor_1-new.png")
	sprites["PH_WARNING"], _, _ = ebitenutil.NewImageFromFile("assets/ph_sensor/ph_sensor_1-new.png")
	sprites["PH_DANGER"], _, _ = ebitenutil.NewImageFromFile("assets/ph_sensor/ph_sensor_1-new.png")

	//tds
	sprites["TDS_NORMAL"], _, _ = ebitenutil.NewImageFromFile("assets/tds_sensor/tds_sensor_1-new.png")
	sprites["TDS_WARNING"], _, _ = ebitenutil.NewImageFromFile("assets/tds_sensor/tds_sensor_1-new.png")
	sprites["TDS_DANGER"], _, _ = ebitenutil.NewImageFromFile("assets/tds_sensor/tds_sensor_1-new.png")

	//temp
	sprites["TEMP_NORMAL"], _, _ = ebitenutil.NewImageFromFile("assets/temp_sensor/temp_sensor_1-new.png")
	sprites["TEMP_WARNING"], _, _ = ebitenutil.NewImageFromFile("assets/temp_sensor/temp_sensor_1-new.png")
	sprites["TEMP_DANGER"], _, _ = ebitenutil.NewImageFromFile("assets/temp_sensor/temp_sensor_1-new.png")

	//ntu
	sprites["NTU_NORMAL"], _, _ = ebitenutil.NewImageFromFile("assets/ntu_sensor/ntu_sensor_1-new.png")
	sprites["NTU_WARNING"], _, _ = ebitenutil.NewImageFromFile("assets/ntu_sensor/ntu_sensor_1-new.png")
	sprites["NTU_DANGER"], _, _ = ebitenutil.NewImageFromFile("assets/ntu_sensor/ntu_sensor_1-new.png")



	//ANIMACIONES DE FLECHAS
	

		// TEMP
	tempFrames := []*ebiten.Image{
		MustNewImageFromFile("assets/temp_sensor/arrowA1.png"),
		MustNewImageFromFile("assets/temp_sensor/arrowA2.png"),
		MustNewImageFromFile("assets/temp_sensor/arrowA3.png"),
	}

	
	// phFrames := []*ebiten.Image{
	// 	MustNewImageFromFile("assets/ph_sensor/arrow_ph1.png"),
	// 	MustNewImageFromFile("assets/ph_sensor/arrow_ph2.png"),
	// 	MustNewImageFromFile("assets/ph_sensor/arrow_ph3.png"),
	// }
	
	// // TDS
	// tdsFrames := []*ebiten.Image{
	// 	MustNewImageFromFile("assets/tds_sensor/arrow_tds1.png"),
	// 	MustNewImageFromFile("assets/tds_sensor/arrow_tds2.png"),
	// 	MustNewImageFromFile("assets/tds_sensor/arrow_tds3.png"),
	// }

	
	// // NTU
	// ntuFrames := []*ebiten.Image{
	// 	MustNewImageFromFile("assets/ntu_sensor/arrow_ntu1.png"),
	// 	MustNewImageFromFile("assets/ntu_sensor/arrow_ntu2.png"),
	// 	MustNewImageFromFile("assets/ntu_sensor/arrow_ntu3.png"),
	// }
	
	
	
	return &GameManager{
		value_ph:   0,
		value_ntu:  0,
		value_temp: 0,
		value_tds:  0,

		 
		iotImg: deviceImg,
		sensorSprites: sprites,
		waterBucketImg: bucketImg,

		tempAnim: NewAnimation(tempFrames, 10),
		// phAnim: NewAnimation(phFrames, 10),
		// tdsAnim: NewAnimation(tdsFrames, 10),
		// ntuAnim: NewAnimation(ntuFrames, 10),

		resolution_w:  1280, 
        resolution_h:  720,
		cancelContext: cancelCtx,
		finish: false,
		exitStatus: exitStatus,
	}
}

func (g *GameManager) Update() error {

	g.tempAnim.Update()
    // g.phAnim.Update()
	// g.tdsAnim.Update()
	// g.ntuAnim.Update()

	// Al presionar Esc o finalizar desde otra goroutine, termina juego
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || g.finish {
		println("Cancelando contexto y cerrando juego...")
		// Cierra todas las goroutines
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


	//PH
	var phSprite *ebiten.Image
    switch currentData.PHState {
    case "WARNING":
        phSprite = g.sensorSprites["PH_WARNING"]
    case "DANGER":
        phSprite = g.sensorSprites["PH_DANGER"]
    default:
        phSprite = g.sensorSprites["PH_NORMAL"]
    }

	if phSprite != nil {
        opPH := &ebiten.DrawImageOptions{}
		opPH.GeoM.Translate(float64(768), float64(428))
        screen.DrawImage(phSprite, opPH)
    }

	//TDS
	var tdsSprite *ebiten.Image
	switch currentData.TDSState {
	case "WARNING":
		tdsSprite = g.sensorSprites["TDS_WARNING"]
	case "DANGER":
		tdsSprite = g.sensorSprites["TDS_DANGER"]
	default:
		tdsSprite = g.sensorSprites["TDS_NORMAL"]
	}

	if tdsSprite != nil {
		opTDS := &ebiten.DrawImageOptions{}
		opTDS.GeoM.Translate(float64(272), float64(188)) 
		screen.DrawImage(tdsSprite, opTDS)
	}

	//TEMPERATURA
	var tempSprite *ebiten.Image
	switch currentData.TempState {
	case "WARNING":
		tempSprite = g.sensorSprites["TEMP_WARNING"]
	case "DANGER":
		tempSprite = g.sensorSprites["TEMP_DANGER"]
	default:
		tempSprite = g.sensorSprites["TEMP_NORMAL"]
	}

	if tempSprite != nil {
		opTemp := &ebiten.DrawImageOptions{}
		opTemp.GeoM.Translate(float64(777), float64(188)) 
		screen.DrawImage(tempSprite, opTemp)
	}

	//NTU
	var ntuSprite *ebiten.Image
	switch currentData.TurbState {
	case "WARNING":
		ntuSprite = g.sensorSprites["NTU_WARNING"]
	case "DANGER":
		ntuSprite = g.sensorSprites["NTU_DANGER"]
	default:
		ntuSprite = g.sensorSprites["NTU_NORMAL"]
	}

	if ntuSprite != nil {
		opNTU := &ebiten.DrawImageOptions{}
		opNTU.GeoM.Translate(float64(272), float64(428)) 
		screen.DrawImage(ntuSprite, opNTU)
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
	g.tempAnim.Draw(screen, float64(700), float64(260))
	// g.phAnim.Draw(screen, float64(784), float64(338)) 
	// g.tdsAnim.Draw(screen, float64(280), float64(100))
	// g.ntuAnim.Draw(screen, float64(280), float64(338))

}


func MustNewImageFromFile(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatalf("Error al cargar imagen (Must): %s - %v", path, err)
	}
	return img
}