package main

import "time"

func main() {
	mine := []string{"rock", "gold", "rock", "rock", "gold", "gold", "rock", "rock", "gold"}
	oreChannel := make(chan string)
	meltChannel := make(chan string)
	doneChannel := make(chan string)

	go finder(mine, oreChannel, doneChannel)
	go breaker(oreChannel, meltChannel)
	go smelter(meltChannel)

	println(<-doneChannel) // waiting all gold to be found
}

func finder(mine []string, oreChannel chan string, doneChannel chan string) {
	for _, ore := range mine {
		if ore != "rock" {
			println("<< found >>>")
			time.Sleep(1000 * time.Millisecond)
			println("<<< sending to break >>>")
			time.Sleep(1000 * time.Millisecond)
			oreChannel <- ore
		}
	}
	doneChannel <- "I'm done here!"
}

func breaker(oreChannel chan string, meltChannel chan string) {
	for ore := range oreChannel {
		println("<<< breaking >>>")
		time.Sleep(1000 * time.Millisecond)
		println("<<< sending to melt >>>")
		time.Sleep(1000 * time.Millisecond)
		meltChannel <- ore
	}
}

func smelter(meltChannel chan string) {
	for range meltChannel {
		println("<<< melting >>>")
		time.Sleep(1000 * time.Millisecond)
	}
}
