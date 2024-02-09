package main

import "fmt"

// This function simply returns an empty struct.
func empty_flags_struct() Flags {
	return Flags {
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	}
}

// This function takes a list of strings (hopefully flags)
// and stores them into a struct for ease of usage. 
func create_flags(flags []string) Flags {
	var result Flags = empty_flags_struct()

	// iterating over the flags passed into the function
	for _, flag := range flags {
		switch flag {
			// atomic flags
			case F_ESCAPE:
				fallthrough
			case F_ESCAPE_LONG:
				result.ESCAPE = true
			case F_NUMBERED:
				fallthrough
			case F_NUMBERED_LONG:
				result.NUMBERED = true
			case F_TRUNCATE:
				fallthrough
			case F_TRUNCATE_LONG:
				result.TRUNCATE = true
			case F_SHOWEND:
				fallthrough
			case F_SHOWEND_LONG:
				result.SHOWEND = true
			case F_SHOWTABS:
				fallthrough
			case F_SHOWTABS_LONG:
				result.SHOWTABS = true
			case F_HELP:
				fallthrough
			case F_HELP_LONG:
				result.HELP = true
			case F_VERSION:
				fallthrough
			case F_VERSION_LONG:
				result.VERSION = true

			// molecule flags
			case F_SHOWALL:
				fallthrough
			case F_SHOWALL_LONG:
				result.ESCAPE = true
				result.SHOWEND = true
				result.SHOWTABS = true
			case F_TABALL:
				result.ESCAPE = true
				result.SHOWTABS = true
			default:
				// default case, shouldn't affect execution of program
				fmt.Printf("unknown flag: %s\n", flag)
		}
	}

	// returning the result
	return result
}

