package application

type SensorDisplayDTO struct {
	PHValue     float32
	// para mostar si esta bueno, maolo etc etc asi con todos
	PHState     string 

	TDSValue    float32
	TDSState    string

	TempValue   float32
	TempState   string

	TurbValue   float32 
	TurbState   string

	IsWaterSafe bool
}
//la que se llama desde la ui papa
type UIStatePusher interface {
	UpdateState(dto SensorDisplayDTO)
}