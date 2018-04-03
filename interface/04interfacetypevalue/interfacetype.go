package main

import "fmt"

type I interface {
	method1()
}

type T struct {
	abc string
}

func (T) method1() {}

func main() {
	var i I = T{"ab"}
	fmt.Println(i)
}
