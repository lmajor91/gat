package main

import (
	"fmt";
	"os";
)

func main() {
	args := os.Args[1:]

	// exiting if nothing was passed
	if len(args) < 1 {
		os.Exit(0)
	}

	// picking out flags from files
	var ins Instructions = parse_instructions(args)

	// debug
	fmt.Println(ins)

	// running main function
	var return_codes = []int{}
	for _, file := range ins.files {
		return_codes = append(return_codes, gat(file, ins.flags))
	}

	var errors int
	for _, code := range return_codes {
		if code == FILE_READ_ERROR {
			errors += 1
		}
	}

	// calculating the proper return code 
	switch errors {
	case 0:
		os.Exit(NORMAL)
	case 1:
		os.Exit(FILE_READ_ERROR)
	default:
		os.Exit(NOT_ALL_FILES_READ)
	}
}

