package main

import "fmt"

// It is possible to receive many parameters 
// with "Variadic functions", just like varargs in Java
func sumAll(numbers ...int) {
	fmt.Print(numbers, " = ")
	total := 0
	for _, n := range numbers {
		total += n
	}

	fmt.Println(total)
}

func main() {
	sumAll(1, 2, 3)
	sumAll(1, 1, 1, 2, 2, 2)

	numbers := []int{7, 7, 7}
	sumAll(numbers...)
}