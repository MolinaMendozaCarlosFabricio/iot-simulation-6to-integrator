package application

// import "simulator.iot.integrator.6th/src/domain/models"
type SensorDisplayDTO struct {
	PHValue     float64
	PHState     string // para mostar si esta bueno, maolo etc etc asi con todos

	TDSValue    float64
	TDSState    string

	TempValue   float64
	TempState   string

	TurbValue   float64
	TurbState   string

	IsWaterSafe bool
}


type SensorReader interface {
	GetLatestReadings() SensorDisplayDTO
}