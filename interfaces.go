package main

import "fmt"

// I interface is a simple contract
type I interface {
	M()
}

// T is a struct which has a property S which is a string
type T struct {
	S string
}

// M is the implementation of I.M() but
// implementing inside T
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	// var i I = T{"hello! :)"}
	// i.M()

	theEmptyInterface()
}

func theEmptyInterface() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
