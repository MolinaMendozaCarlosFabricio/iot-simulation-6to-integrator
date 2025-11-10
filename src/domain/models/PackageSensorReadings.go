package models

type PackageSensorReadings struct {
	IdUser         string                `json:"idUser"`
	IdFilter      string                 `json:"idFilter"`
	SensorReadings []SensorReadingChunck `json:"sensorReadings"`
}

type SensorReadingChunck struct {
	Id			int		`json:"id"`
	Value 		float32	`json:"value"`
	Date		string	`json:"date"`
	SensorId	string	`json:"sensor_id"`
}