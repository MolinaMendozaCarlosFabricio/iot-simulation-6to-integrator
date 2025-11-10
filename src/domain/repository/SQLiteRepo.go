package repository

import "simulator.iot.integrator.6th/src/domain/models"

type SQLiteRepo interface {
	CreateSensorReading(value float64, idSensor string, backed bool) error
	GetSensorReadingsNotSent() ([]models.SensorReadings, error)
	MarkSensorReadingSent(id int) error
	CloseDB()
}