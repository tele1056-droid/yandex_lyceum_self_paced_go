package main

import (
	"fmt"
)

const pricePerKm float64 = 10.0
const pricePerMinute float64 = 2.0

type TripParameters struct {
	Distance float64
	Duration float64
}

func CalculateBasePrice (t TripParameters) float64 {
	return t.Distance * pricePerKm + t.Duration * pricePerMinute 
}

func main() {
	//test := TripParameters{Distance: 62.3, Duration: 12.0}

	trip1 := TripParameters{Distance: 15.0, Duration: 22.3}
	fmt.Println(CalculateBasePrice(trip1))
}
	
