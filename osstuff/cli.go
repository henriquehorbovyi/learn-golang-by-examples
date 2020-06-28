package osstuff

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// CLIArguments is
func CLIArguments() {
	arguments := os.Args
	someArguments := os.Args[2:]
	firstArgument := os.Args[1]

	fmt.Println("First argument >", firstArgument)
	fmt.Println("Some arguments >", someArguments)
	fmt.Println("All arguments >", arguments)
}

// TaskProgress mocks a CLI task loading progress
func TaskProgress() {
	const total = 30
	percent := 0

	bar := ""
	for i := 0; i < total; i++ {
		percent = i * 100 / total
		fmt.Printf("[%v%v](%v%v)", bar, strings.Repeat("=", i)+">", percent, "%")
		time.Sleep(300 * time.Millisecond)
		bar = fmt.Sprintf("%v", strings.Repeat("\b", i+7))
	}
	fmt.Printf("[%v%v%v]", bar, strings.Repeat("=", total), " Done!")
	fmt.Println()
}
