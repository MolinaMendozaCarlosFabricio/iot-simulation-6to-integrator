package models

type SensorReadings struct {
	IdSensorReading	int
	Value 			float32
	timestamp		string
	IdSensor		string
	backed			bool
}