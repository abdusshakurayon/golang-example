package main

import "fmt"

type Bootcamp struct {
	Lat float64
	Lon float64
}

func main() {
	x := new(Bootcamp) //gives a pointer
	y := &Bootcamp{}
	fmt.Println(*x == *y)
}
