package functions

import (
	"context"
	"log"
	"time"

	"simulator.iot.integrator.6th/src/application"
	"simulator.iot.integrator.6th/src/domain/models"
)

func ExecuteSensors(
	msruc 			application.ManageSensorReadingsUC, 
	vruc 			application.ValidateReadingsUC,
	sensorFunctions []func()(float32, error),
	ctx 			context.Context,
	configDevice 	models.ConfigDevice,
) {
	chanOut := make(chan models.FanOutData)
	chanIn := make(chan models.FanInData)
	n_iterations := 0
	for {
		for range sensorFunctions {
			go sensorWorker(chanOut, chanIn)
		}

		for i := 0; i < 4; i++ {
			chanOut <- models.FanOutData{
				SensorID: configDevice.Sensors_info[i].Sensor_id,
				Name: configDevice.Sensors_info[i].Sensor,
				Function: sensorFunctions[i],
			}
		}

		var recordsTaken []models.SensorReadingsTaken
		for i := 0; i < 4; i++ {
			result := <- chanIn
			if result.Err != nil {
				log.Fatal(result.Err)
			}
			recordsTaken = append(recordsTaken, result.SensorReading)
		}

		var sendThisRecords models.PackageSensorReadings
		sendThisRecords.IdFiltrer = configDevice.Device_id
		sendThisRecords.IdUser = configDevice.User_id
		for i := 0; i < 4; i++ {
			sendThisRecords.SensorReadings = append(
				sendThisRecords.SensorReadings,
				models.SensorReadingChunck{
					Id: 0,
					SensorId: "",
					Date: "",
					Value: 0,
				},
			)
		}

		println("Iteración:", n_iterations)
		for _, record := range recordsTaken {
			sensorReadingChunk := models.SensorReadingChunck{
				Id: record.MeasurementId,
				SensorId: record.SensorId,
				Date: record.ReadingDate,
				Value: record.Value,
			}
			println(record.Name)
			println(record.MeasurementId)
			println(record.SensorId)
			println(record.ReadingDate)
			println(record.Value)
			switch record.Name {
				case "Sensor de Temperatura":
					sendThisRecords.SensorReadings[0] = sensorReadingChunk
				case "Sensor de TDS":
					sendThisRecords.SensorReadings[1] = sensorReadingChunk
				case "Sensor de PH":
					sendThisRecords.SensorReadings[2] = sensorReadingChunk
				case "Sensor de Turbidez":
					sendThisRecords.SensorReadings[3] = sensorReadingChunk
			}
		}

		n_iterations++

		select{
			case <- ctx.Done():
				return
			default:
				println("Continuando toma de datos")
				time.Sleep(5 * time.Second)
		}
	}
}