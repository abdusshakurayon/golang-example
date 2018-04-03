package main

import (
	"fmt"
	"golang-example/interface/06dependencyinjection/helper"
)

func main() {
	u := &helper.User{"Matt", "Aimonetti"}
	fmt.Println(helper.Greet(u))
}
