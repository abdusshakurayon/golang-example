package main

import "fmt"

type User struct {
	FirstName, LastName string
}

func (u *User) Name() string {
	return fmt.Sprintf("User %s %s", u.FirstName, u.LastName)
}

type Customer struct {
	Id       int
	FullName string
}

func (c *Customer) Name() string {
	return fmt.Sprintf("Customer name %s and id %d", c.FullName, c.Id)
}

func (c *Customer) Address() string {
	return fmt.Sprintf("Customer name is address %s ", c.FullName)
}

type Namer interface {
	Name() string
}

type Address interface {
	Namer
	Address() string
}

func Greet(n Address) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

func main() {
	u := &User{"Matt", "Aimonetti"}
	fmt.Println(Greet(u))
	//Will not compile, User does not implements address
	c := &Customer{42, "Francesc"}
	fmt.Println(Greet(c))

}
