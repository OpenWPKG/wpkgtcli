//go:build !windows
// +build !windows

package main

import (
	"fmt"
)

const (
	NEWLINE = "\n"

	color_Reset  = "\033[0m"
	color_Black  = "\033[30m"
	color_Red    = "\033[31m"
	color_Green  = "\033[32m"
	color_Yellow = "\033[33m"
	color_Blue   = "\033[34m"
	color_Purple = "\033[35m"
	color_Cyan   = "\033[36m"
	color_Gray   = "\033[37m"
	color_White  = "\033[97m"
)



func mOk(message ...any) {
	fmt.Print("[", color_Green, "ok", color_Reset, "] ")
	fmt.Println(message...)
}
func mInfo(message ...any) {
	fmt.Print("[", color_Yellow, "info", color_Reset, "] ")
	fmt.Println(message...)
}
func mError(message ...any) {
	fmt.Print("[", color_Red, "error", color_Reset, "] ")
	fmt.Println(message...)
}
func mDebug(message ...any) {
	if !DEBUG {
		return
	}
	fmt.Print("[", color_Blue, "debug", color_Reset, "] ")
	fmt.Println(message...)
}