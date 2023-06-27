package pglmmylog

import (
	"fmt"
	"time"
)

type debug_mode int

const version = "1.0.0.01"

const (
	MODE_NONE      = iota // Do nothing
	MODE_DEFAULT          // Default level for debug
	MODE_DEEP_DIVE        // Debug higher level in loop.
)

const CURRENT_DEBUG_MODE debug_mode = MODE_DEFAULT
const isINFO_MODE = true
const FORMAT_DEBUG_DATETIME string = "2006-01-02 15:04:05.0000"

func GetVersion() string {
	return version
}

// Write log to console
func WriteInfo(a ...any) {
	fmt.Println(time.Now().Format(FORMAT_DEBUG_DATETIME), a)
}

// print when is debug mode enabled
func PrintDebug(mode debug_mode, a ...any) {
	if CURRENT_DEBUG_MODE <= MODE_NONE {
		return // do nothing
	}

	if CURRENT_DEBUG_MODE >= mode {
		fmt.Print(time.Now().Format(FORMAT_DEBUG_DATETIME) + " ")
		fmt.Println(a...)
	}
}

// print information
func PrintInfo(a ...any) {
	if isINFO_MODE {
		fmt.Print(time.Now().Format(FORMAT_DEBUG_DATETIME) + " ")
		fmt.Println(a...)
	}
}
