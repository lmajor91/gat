package main

import (
	"fmt";
	"os";
)

func main() {
	args := os.Args[1:]

	fmt.Println(args);

	// exiting if nothing was passed
	if len(args) < 1 {
		os.Exit(0)
	}

	gat(args[0])

	// exiting from the program
	os.Exit(0)
}

