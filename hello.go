package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

type number int

func (n number) isPositive() bool {
	return n > 0
}

func main() {

	// var age number = -42
	// fmt.Println(age.isPositive())

	// b := make([]int, 2)
	// fmt.Println(b)

	// sum := add(21, 21)
	// fmt.Println("Sum is", sum)

	// a, b := swap("world", "hello")
	// fmt.Println(a, b)

	// x, y := split(2)
	// fmt.Println(x, y)

	// variables()

	// loops()

	// ifStatements()

	// switchSamples()

	// deferSamples()
}

func helloWorld() {
	fmt.Printf("Hello World\n")
}

func mathStuff() {
	fmt.Println("That's a random number: ", rand.Intn(10))
	fmt.Println("PI is: ", math.Pi)
}

func whatTimeIsIt() {
	fmt.Println("Current time: ", time.Now())
}

// When you have two consecutive variables with the same type
// into your function you can write the type only in the last variable
func add(x int, y int) int {
	return x + y
}

// in Go you can have many returns
func swap(s1, s2 string) (string, string) {
	return s2, s1
}

// returns in Go can be named :O
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func variables() {
	// var defines a variable that can be assigned later
	var x int
	fmt.Println(x) // default value for an int unassigned is 0 :D

	x = 2
	fmt.Println(x)
	x = 1
	fmt.Println(x)
	x = 0
	fmt.Println(x)

	// Inside a function, the := short
	// assignment statement can be used in
	// place of a var declaration with implicit type.
	y := 21
	y = 42
	fmt.Println(y)

	// You can create variables inline
	var c, python, java = true, false, "no! :D"
	fmt.Println(c, python, java)

	var (
		isAvailable bool       = false
		z           complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T | Value: %v\n", isAvailable, isAvailable)
	fmt.Printf("Type: %T | Value: %v\n", z, z)

	const (
		Big = 1 << 2
	)
	a := float32(Big)
	fmt.Println(a)
}

func loops() {
	// for
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum is:", sum)

	// "while"
	for sum < 0 {
		sum = sum - 1
	}
	fmt.Println("Now sum is:", sum)

	// Infinity loop
	// for { }
}

func ifStatements() {
	age := 17
	if age >= 18 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func switchSamples() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Print("OS X\n")
	case "linux":
		fmt.Print("Linux\n")
	default:
		fmt.Print(os, "\n")
	}

	fmt.Println("\nWhen's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// Switch with conditions
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferSamples() {
	fmt.Println("counting...")

	i := 0
	for i < 10 {
		defer fmt.Println(i)
		i++
	}

	fmt.Println("done")
}
