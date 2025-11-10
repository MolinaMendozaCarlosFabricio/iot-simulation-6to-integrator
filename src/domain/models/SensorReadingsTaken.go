package models

type SensorReadingsTaken struct {
	SensorId		string
	MeasurementId 	int
	Name 			string
	Value			float32
	ReadingDate		string
}