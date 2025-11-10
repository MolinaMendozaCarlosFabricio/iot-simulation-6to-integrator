package models

type PackageSensorReadings struct {
	IdUser			string					`json:"id_user"`
	IdFiltrer		string					`json:"id_device"`
	SensorReadings []SensorReadingChunck	`json:"sensorReadings"`
}

type SensorReadingChunck struct {
	Id			int		`json:"id"`
	Value 		float32	`json:"value"`
	Date		string	`json:"date"`
	SensorId	string	`json:"sensor_id"`
}