package main

import (
	"bytes"
	"fmt"
	"os"
)

func parse_instructions(arguments []string) (Instructions) {
	_flags := []string{} 
	_files := []string{} 

	// iterating over the arguments
	for i := 0; i < len(arguments); i++ {
		var _s = arguments[i];

		// skip if length is zero
		if len(_s) == 0 {
			continue
		}

		// checking if it's a flag e.g., does it start with '-'?
		if _s[0] != '-' || _s == "-" {
			// input was not a flag, was either stdin or a file
			_files = append(_files, _s)
		} else {
			// input was a flag

			// checking for multi flag e.g., -vET
			if len(_s) > 2 && _s[1] != '-' {
				// iterating over the flag string
				for j := 1; j < len(_s); j++ {
					var b bytes.Buffer
					b.WriteByte('-')
					b.WriteByte(_s[j])
					_flags = append(_flags, b.String())
				}
			} else {
				// normal two char flag
				_flags = append(_flags, _s)
			}
		}
	}

	return Instructions{ flags: create_flags(_flags), files:  _files }
}

func gat(file string, flags Flags) (int) {
	// reading the file from the command line
	fd, err := os.ReadFile(file)

	// error reading the file
	if err != nil {
		fmt.Printf("Couldn't read file: %s\n", file)
		return FILE_READ_ERROR
	}

	// printing the file
	for i := 0; i < len(fd); i++ {
		fmt.Printf("%c", fd[i])
	}

	return NORMAL
}

