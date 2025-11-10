package models

type ConfigDevice struct {
	Device_id		string	`json:"device_id"`
	User_id			string	`json:"user_id"`
	Model			string	`json:"model"`
	Created_at		string	`json:"created_at"`
	Sensors_info	[]ConfigSensor `json:"sensors_info"`
}

type ConfigSensor struct {
	Sensor				string	`json:"sensor"`
	Sensor_name_model	string	`json:"sensor_name_model"`
	Sensor_id			string	`json:"sensor_id"`
}