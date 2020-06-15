package main

import "fmt"

type person struct {
	name string
	age  int
}

func (person person) isOlderThen(age int) bool {
	return person.age > age
}

func main() {
	var me = person{
		name: "Henrique",
		age:  23,
	}

	if me.isOlderThen(18) {
		fmt.Println("You can drive!")
	} else {
		fmt.Println("Can't drive dude!")
	}
}
