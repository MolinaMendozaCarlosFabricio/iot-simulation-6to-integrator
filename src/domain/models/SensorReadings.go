package models

type SensorReadings struct {
	IdSensorReading	int
	Value 			float32
	Timestamp		string
	IdSensor		string
	Backed			bool
}