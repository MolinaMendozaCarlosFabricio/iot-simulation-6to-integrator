package functions

import (
	"time"

	"simulator.iot.integrator.6th/src/domain/models"
)

func sensorWorker(
	chanOut chan models.FanOutData,
	chanIn chan models.FanInData,
) {
	structData := <- chanOut
	measurement, err := structData.Function()
	record := models.SensorReadingsTaken{
		SensorId: structData.SensorID,
		Name: structData.Name,
		MeasurementId: 0,
		Value: measurement,
		ReadingDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	
	chanIn <- models.FanInData{SensorReading: record, Err: err}
}