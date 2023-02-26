package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)



func input(prompt string, defaultOutput string) string {
	fmt.Printf("%s [%s]: ", prompt, defaultOutput)
	_in, err := Input.ReadString('\n')
	line := strings.Trim(_in, NEWLINE)
	mDebug([]byte(_in), "\""+line+"\"")
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
	fmt.Printf("%s [%s]: ", prompt, fmt.Sprint(ternary(defaultOutput == "no", defaultOutput, defaultOutput_print)))
	_in, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		mError("Error reading value:", err.Error())
		os.Exit(1)
	}
	line := strings.Trim(string(_in), NEWLINE)
	if line == "" {
		return defaultOutput
	} else {
		return line
	}
}
