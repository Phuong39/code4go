package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Printf("Please enter some input: ")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Printf("The input was: %s", input)
	}
}
