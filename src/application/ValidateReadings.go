package application

import "simulator.iot.integrator.6th/src/domain/models"

type ValidateReadingsUC struct{}

func NewValidateReadingsUC() *ValidateReadingsUC {
	return &ValidateReadingsUC{}
}

func (uc *ValidateReadingsUC) Execute(p models.PackageSensorReadings) bool {
	if p.SensorReadings[0].Value < 10 || p.SensorReadings[0].Value > 40 { return false }
	if p.SensorReadings[2].Value < 5 || p.SensorReadings[2].Value > 9 { return false }
	if p.SensorReadings[3].Value > 40 { return false }
	if p.SensorReadings[1].Value > 1500 && p.SensorReadings[1].Value > 40 { return false }
	return true

}