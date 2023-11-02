package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mycliapp <command>")
		return
	}

	command := os.Args[1]
	switch command {
	case "hello":
		fmt.Println("Hello, World!")
	default:
		fmt.Println("Unknown command:", command)
	}
}
