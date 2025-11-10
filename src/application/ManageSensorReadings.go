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
		return fmt.Errorf("Error al mandar mensaje al WebSocket por AMQP:", err)
	}

	// Mandar a la cola de la API de sensores
	err = uc.RabbitMQRepo.PublishMessage(".measurements", pack.SensorReadings)
	if err != nil {
		return fmt.Errorf("Error al mandar mensaje a la API por AMQP:", err)
	}

	// Agregar respaldo en una base de datos local

	return nil
}