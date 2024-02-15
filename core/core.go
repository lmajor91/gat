package main

import (
	"bytes"
	"fmt"
	"os"
	"math"
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

	// numbering the lines
	var line_number int = 0
	var line_fmt_str string = ""

	// if NUMBERED is true, count the new lines to nicely format
	// everything
	if flags.NUMBERED {
		for i := 0; i < len(fd); i++ {
			if fd[i] == '\n' {
				line_number += 1
			}
		}

		// accounting for either no new lines
		// or the final line
		line_number += 1

		// calculating padding
		line_fmt_str = fmt.Sprintf("%%%dd: ", int(math.Log10(float64(line_number))))

		// resetting line_number for other use
		line_number = 1
	}

	// printing initial line number
	if flags.NUMBERED {
		fmt.Printf(line_fmt_str, line_number)
		line_number++
	}

	// printing the file
	for i := 0; i < len(fd); i++ {
		// flags that trigger based on a newline character
		if fd[i] == '\n' {
			// showing end of lines
			if flags.SHOWEND {
				fmt.Print("$")
			}
		}

		// printing the char
		fmt.Printf("%c", fd[i])
	
		// flags that trigger based on a newline character
		if fd[i] == '\n' {
			if flags.NUMBERED {
				fmt.Printf(line_fmt_str, line_number)
				line_number++
			}
		}
	}

	return NORMAL
}

