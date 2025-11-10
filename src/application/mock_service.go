package application

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type MockSensorService struct {
	mutex      sync.RWMutex
	latestData SensorDisplayDTO
}

func NewMockSensorService() *MockSensorService {
	svc := &MockSensorService{}
	go svc.runMockDataGenerator()
	return svc
}

func (s *MockSensorService) GetLatestReadings() SensorDisplayDTO {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.latestData
}

func (s *MockSensorService) runMockDataGenerator() {
	log.Println("iniciando sumilador")
	ticker := time.NewTicker(500 * time.Millisecond) 
	defer ticker.Stop()

	for {
		<-ticker.C

		// simulaoooo
		tempVal := 10 + rand.Float64()*30 
		phVal := 5 + rand.Float64()*4     
		turbVal := rand.Float64() * 50    
		tdsVal := rand.Float64() * 2000  
		
		tempState := "NORMAL"
		if tempVal < 15 || tempVal > 35 {
			tempState = "WARNING"
		}
		if tempVal < 10 || tempVal > 40 {
			tempState = "DANGER"
		}

		phState := "NORMAL"
		if phVal < 6.0 || phVal > 8.0 {
			phState = "WARNING"
		}
		if phVal < 5.0 || phVal > 9.0 {
			phState = "DANGER"
		}
        
       
        turbState := "NORMAL"
        tdsState := "NORMAL"


		
		isSafe := true
		if tempVal < 10 || tempVal > 40 {
			isSafe = false
		}
		if phVal < 5 || phVal > 9 {
			isSafe = false
		}
		if turbVal > 40 {
			isSafe = false
		}
		if tdsVal > 1500 && turbVal > 40 {
			isSafe = false
		}

		
		data := SensorDisplayDTO{
			PHValue:     phVal,
			PHState:     phState,
			TDSValue:    tdsVal,
			TDSState:    tdsState,
			TempValue:   tempVal,
			TempState:   tempState,
			TurbValue:   turbVal,
			TurbState:   turbState,
			IsWaterSafe: isSafe,
		}

		
		s.mutex.Lock() 
		s.latestData = data
		s.mutex.Unlock() 
	}
}