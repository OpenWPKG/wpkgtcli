package main

import (
	"fmt"
	"strings"
	"os"

	"github.com/TwiN/go-color"
	"golang.org/x/term"
)

func mOk(message ...any) {
	fmt.Print("[", color.Green, "ok", color.Reset, "] ")
	fmt.Println(message...)
}
func mInfo(message ...any) {
	fmt.Print("[", color.Yellow, "info", color.Reset, "] ")
	fmt.Println(message...)
}
func mError(message ...any) {
	fmt.Print("[", color.Red, "error", color.Reset, "] ")
	fmt.Println(message...)
}
func mDebug(message ...any) {
	if !DEBUG {return}
	fmt.Print("[", color.Blue, "debug", color.Reset, "] ")
	fmt.Println(message...)
}

func input(prompt string, defaultOutput string) string {
	fmt.Printf("%s [%s]: ", prompt, defaultOutput)
	_in, err := Input.ReadString('\n')
	line := strings.Split(_in, "\n")[0]
	if err != nil {
		mError("Error reading value:", err.Error())
		os.Exit(1)
	}
	if line == "" {
		return defaultOutput
	} else {
		return line
	}
}

func input_p(prompt string, defaultOutput string) string {
	defaultOutput_print := ""
	for i := 0; i < len(defaultOutput); i++ {
		defaultOutput_print += "*"
	}
	fmt.Printf("%s [%s]: ", prompt, fmt.Sprint(ternary(defaultOutput=="no", defaultOutput, defaultOutput_print)))
	_in, err := term.ReadPassword(0)
	fmt.Println()
	line := strings.Split(string(_in), "\n")[0]
	if err != nil {
		mError("Error reading value:", err.Error())
		os.Exit(1)
	}
	if line == "" {
		return defaultOutput
	} else {
		return line
	}
}