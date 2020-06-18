package osstuff

import (
	"fmt"
	"os"
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
