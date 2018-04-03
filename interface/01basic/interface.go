package main

import "fmt"

type User struct {
	FirstName, LastName string
}

/**
An interface type is defined by a set of methods. A value of interface type can hold any value that implements those methods.
*/

type Namer interface {
	Name() string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

func main() {
	u := &User{"Matt", "Aimonetti"}
	fmt.Println(Greet(u))
}
