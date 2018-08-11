package snek

import "github.com/amy911/amy911/syspath"

var (
	// Set this in an `init()` function somewhere
	DefaultConfigType string = "yaml"

	// Initialize this in an `init()` function somewhere
	SysPath *syspath.SysPath
)
