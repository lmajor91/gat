package main

// the Flags struct, this determines how the program executes
type Flags struct {
	ESCAPE		bool;	// escape all escape characters like \n and nullbyte
	NUMBERED	bool;	// number all lines
	TRUNCATE	bool;	// truncate repeated empty lines
	SHOWEND		bool;	// show the end of lines
	SHOWTABS	bool;	// show tabs as another character
	HELP		bool;	// display help message
	VERSION		bool;	// show version info
}

// this was made so that the `parse_instruction` function
// only has to iterate over the command line arguments once.
// this struct should really only be passed into `gat`
type Instructions struct {
	flags	Flags;
	files	[]string;
}

