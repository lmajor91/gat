package main

// errors
const NORMAL int = 0
const FILE_READ_ERROR int = 1
const NOT_ALL_FILES_READ int = 2

// atoms, flags that do a single thing
const F_ESCAPE			string = "-v"
const F_ESCAPE_LONG 	string = "--show-nonprinting"
const F_NUMBERED		string = "-n"
const F_NUMBERED_LONG	string = "--number"
const F_TRUNCATE		string = "-s"
const F_TRUNCATE_LONG	string = "--squeeze-blank"
const F_SHOWEND			string = "-E"
const F_SHOWEND_LONG	string = "--show-ends"
const F_SHOWTABS		string = "-T"
const F_SHOWTABS_LONG	string = "--show-tabs"
const F_HELP			string = "-h"
const F_HELP_LONG		string = "--help"
const F_VERSION			string = "-V"
const F_VERSION_LONG	string = "--version"

// molecules, flags that do more than one thing
const F_SHOWALL			string = "-A"
const F_SHOWALL_LONG	string = "--show-all"
const F_TABALL			string = "-t"

