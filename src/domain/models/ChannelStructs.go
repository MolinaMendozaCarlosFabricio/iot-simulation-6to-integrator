package models

type FanOutData struct {
	SensorID 	string
	Name 		string
	Function 	func()(float32, error)
}

type FanInData struct {
	SensorReading 	SensorReadingsTaken
	Err				error
}