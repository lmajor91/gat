package main

import (
	"fmt";
	"os";
)

func gat(filename string) (int) {
	// reading the file from the command line
	file, err := os.ReadFile(filename)

	// error reading the file
	if err != nil {
		fmt.Printf("Couldn't read file: %s\n", filename)
	}

	// printing the file
	for i := 0; i < len(file)-1; i++ {
		fmt.Printf("%c", file[i])
	}

	return 0
}

