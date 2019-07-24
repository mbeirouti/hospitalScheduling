package main

import (
	"errors"
	"log"
	"time"
)

var (
	totalTime = 40
	timePerAppointment

)

type City struct {
	MildProductionRatio     float64
	ModerateProductionRatio float64
	UrgentProductionRatio   float64
	ProductionFunction      func(*[]Patient)
}

func NewCity(mild, moderate, urgent float64, productionFunction func()) (*City, error) {
	if (mild + moderate + urgent) != 1.00 {
		return nil, errors.New("ratios don't sum to one")
	}
	return &City{
		MildProductionRatio:     mild,
		ModerateProductionRatio: moderate,
		UrgentProductionRatio:   urgent,
		ProductionFunction:      productionFunction,
	}, nil
}

type Hospital struct {
	MildQueue   *[]Patient
	UrgentQueue *[]Patient
	Time 
}

type Patient struct {
	WaitTime time.Time
	Served   bool
}

func main() {

	mild := 0.15
	moderate := 0.55
	urgent := 0.3

	montreal, err := NewCity(mild, moderate, urgent, nil)
	if err != nil {
		log.Fatal("Error creating new city")
	}

	go func() {
		mildMax := mild
		moderateMax := mildMax + moderate
		urgentMax := moderateMax + urgent
		
		diceRoll := rand.Float64()
		if diceRoll >= 0 && diceRoll < mildMax {
			
		} else if diceRoll >= mildMax && diceRoll < moderateMax {

		} else if diceRoll >= moderateMax && diceRoll < urgentMax {

		}

	}()

	montreal
}
