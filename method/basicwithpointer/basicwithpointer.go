package main

import "fmt"

/**
 avoid copying the value on each method call (more efficient if the value type is a large struct)
 */
type user struct {
	name, age string
}

func (u *user) greeting() string {
	return fmt.Sprintf("%s age is %s", u.name, u.age)
}

func main() {
	u := &user{"Monstar", "20"}
	fmt.Println(u.greeting())
}
