package sensors

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"os"

	"simulator.iot.integrator.6th/src/application"
	"simulator.iot.integrator.6th/src/domain/models"
	"simulator.iot.integrator.6th/src/infrastructure/adapters/sensors/functions"
)

type SimulatedSensorsManager struct {
	manageSensorReadingsUc 	application.ManageSensorReadingsUC
	validateReadingsUC 		application.ValidateReadingsUC
	ctx						context.Context
	config 					models.ConfigDevice
}

func NewSensorsManager(
	ctx context.Context,
) *SimulatedSensorsManager {
	return&SimulatedSensorsManager{
		manageSensorReadingsUc: *application.NewManageSensorReadingsUC(),
		validateReadingsUC: *application.NewValidateReadingsUC(),
		ctx: ctx,
		config: *getConfigSensor(),
	}
}

func (sm *SimulatedSensorsManager)InitSensorReadings(n_instances int) {
	var sensorFunctions []func()(float32, error)
	sensorFunctions = append(sensorFunctions, sm.measureTemp)
	sensorFunctions = append(sensorFunctions, sm.measureTDS)
	sensorFunctions = append(sensorFunctions, sm.measurePH)
	sensorFunctions = append(sensorFunctions, sm.measureNTU)
	for i := 1; i <= n_instances; i++ {
		go functions.ExecuteSensors(
			sm.manageSensorReadingsUc,
			sm.validateReadingsUC,
			sensorFunctions,
			sm.ctx,
			sm.config,
		)
	}
}

func getConfigSensor()*models.ConfigDevice{
	var c models.ConfigDevice
	file, err := os.Open("./src/config/userdata.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}
	return &c
}

func (sm *SimulatedSensorsManager)measureTemp() (float32, error) { return 10 + rand.Float32()*30, nil }
func (sm *SimulatedSensorsManager)measurePH() (float32, error) { return 5 + rand.Float32()*4, nil }
func (sm *SimulatedSensorsManager)measureNTU() (float32, error) { return rand.Float32() * 50, nil }
func (sm *SimulatedSensorsManager)measureTDS() (float32, error) { return rand.Float32() * 2000, nil }