package main

import "fmt"

/**
Can not duplicate anonymous field
*/
type Person struct {
	string
	int
}

func main() {
	p := Person{"Monstar", 50}
	fmt.Println(p)
	p.string = "Sekai"
	p.int = 20
	fmt.Println(p)
}
