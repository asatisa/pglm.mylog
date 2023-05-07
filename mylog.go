package pglmmylog

import (
	"fmt"
	"time"
)

type debug_mode int

const version = "1.0.0.0"

const (
	MODE_NONE = iota //do nothing
	MODE_DEFAULT
	MODE_DEEP_DIVE
)

const iDEBUG_MODE debug_mode = MODE_DEFAULT
const isINFO_MODE = true
const FORMAT_TIME string = "2006-01-02 15:04:05.0000"

func GetVersion() string {
	return version
}

// Write log to console
func WriteInfo(a ...any) {
	fmt.Println(time.Now().Format(FORMAT_TIME), a)
}

// print when is debug mode enabled
func PrintDebug(mode debug_mode, a ...any) {
	if iDEBUG_MODE <= MODE_NONE {
		return // do nothing
	}

	fmt.Println(time.Now().Format(FORMAT_TIME), a)
}

// print information
func PrintInfo(a ...any) {
	if isINFO_MODE {
		fmt.Println(time.Now().Format(FORMAT_TIME), a)
	}
}
