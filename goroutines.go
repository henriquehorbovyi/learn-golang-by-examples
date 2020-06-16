package main

import (
	"fmt"
	"time"
)

func main() {
	channelSample()
}

func goroutinesSample() {
	go process("hello")
	process("world")
}

func process(message string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("processing...", message)
	}
}

func channelSample() {
	numbers := []int{1, 1, 1, 2, 2, 2}
	channel := make(chan int)

	go sum(numbers[:len(numbers)/2], channel) // first slice part
	go sum(numbers[len(numbers)/2:], channel) // second slice part
	go sum(numbers[:], channel)               // entire slice

	sumA := <-channel // receive from the channel
	sumB := <-channel // receive from the channel
	sumC := <-channel // receive from the channel

	fmt.Println(sumA)
	fmt.Println(sumB)
	fmt.Println(sumC)
}

func sum(numbers []int, channel chan int) {
	fmt.Println("executing sum...")
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	time.Sleep(3000 * time.Millisecond)
	channel <- sum // send sum of numbers to channel
}
