package main

import (
	"fmt"
	"time"
)

/**
Struct can be thought as a class
*/

type Bootcamp struct {
	Lat  float64
	Lon  float64
	Date time.Time
}

func main() {
	fmt.Println(Bootcamp{
		Lat:  34.012836,
		Lon:  -118.495338,
		Date: time.Now(),
	})
}
