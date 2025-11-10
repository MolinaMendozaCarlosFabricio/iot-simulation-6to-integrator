package ui

import (
	"context"
	"log"

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

	//inyectada
	sensorService application.SensorReader
	latestData application.SensorDisplayDTO
	
	//sprites, fondos, img, etc.
	iotImg *ebiten.Image
	waterBucketImg *ebiten.Image
	sensorSprites map[string]*ebiten.Image
}

func NewGame(svc application.SensorReader) *GameManager {	
	
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


	
	
	
	return &GameManager{
		value_ph:   0,
		value_ntu:  0,
		value_temp: 0,
		value_tds:  0,
		sensorService: svc,
		 
		iotImg: deviceImg,
		sensorSprites: sprites,
		waterBucketImg: bucketImg,

		resolution_w:  1280, 
        resolution_h:  720,
	}
}

func (g *GameManager) Update() error {
	g.latestData = g.sensorService.GetLatestReadings()
    
	return nil
}

func (g *GameManager) Layout(outsideWidth, outsideHeight int) (int, int){
	return g.resolution_w, g.resolution_h
}

func (g *GameManager) Draw(screen *ebiten.Image){
	
	
	screen.Fill(color.White)


	//iot en medio
	if g.iotImg != nil {
		opDevice := &ebiten.DrawImageOptions{}
		imgWidth, imgHeight := g.iotImg.Size()
		centerX := float64(g.resolution_w/2) - float64(imgWidth/2)
		centerY := float64(g.resolution_h/2) - float64(imgHeight/2)
		opDevice.GeoM.Translate(centerX, centerY)
		screen.DrawImage(g.iotImg, opDevice)
	}

	//agua
	if g.waterBucketImg != nil {
		opBucket := &ebiten.DrawImageOptions{}
		imgWidth, imgHeight := g.waterBucketImg.Size()
		centerX := float64(g.resolution_w/2) - float64(imgWidth/2)
		centerY := float64(g.resolution_h/2) - float64(imgHeight/2) + 150
		opBucket.GeoM.Translate(centerX, centerY)
		screen.DrawImage(g.waterBucketImg, opBucket)
	}


	//PH
	var phSprite *ebiten.Image
    switch g.latestData.PHState {
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
	switch g.latestData.TDSState {
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
	switch g.latestData.TempState {
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
	switch g.latestData.TurbState {
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
    if g.latestData.IsWaterSafe {
        safeText = "REUTILIZABLE"
    }

	debugText := fmt.Sprintf(
        "ESTADO: %s\n\nPH: %.2f (%s)\nTEMP: %.2f (%s)\nTDS: %.2f (%s)\nTURB: %.2f (%s)",
        safeText,
        g.latestData.PHValue, g.latestData.PHState,
        g.latestData.TempValue, g.latestData.TempState,
        g.latestData.TDSValue, g.latestData.TDSState,
        g.latestData.TurbValue, g.latestData.TurbState,
    )
    ebitenutil.DebugPrint(screen, debugText)



}