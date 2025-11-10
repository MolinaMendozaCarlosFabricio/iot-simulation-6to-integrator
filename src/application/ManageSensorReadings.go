package application

import (
	"fmt"

	"simulator.iot.integrator.6th/src/domain/models"
	"simulator.iot.integrator.6th/src/domain/repository"
)

type ManageSensorReadingsUC struct {
	RabbitMQRepo repository.RabbitMQRepo
}

func NewManageSensorReadingsUC(rabbitRepo repository.RabbitMQRepo) *ManageSensorReadingsUC {
	return &ManageSensorReadingsUC{
		RabbitMQRepo: rabbitRepo,
	}
}

func (uc *ManageSensorReadingsUC) Execute(pack models.PackageSensorReadings) error {
	// Mandar a la cola del websocket
	err := uc.RabbitMQRepo.PublishMessage("websocket_topic.many_readings", pack)
	if err != nil {
		return fmt.Errorf("Error al mandar mensaje por AMQP:", err)
	}

	return nil
}