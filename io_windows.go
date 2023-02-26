package main
import (
	"fmt"

	"github.com/gookit/color"
)

const (
	NEWLINE = "\r\n"
)


func mOk(message ...any) {
	color.Print("[<green>ok</>] ")
	fmt.Println(message...)
}
func mInfo(message ...any) {
	color.Print("[<yellow>info</>] ")
	fmt.Println(message...)
}
func mError(message ...any) {
	color.Print("[<red>error</>] ")
	fmt.Println(message...)
}
func mDebug(message ...any) {
	if !DEBUG {
		return
	}
	color.Print("[<blue>debug</>] ")
	fmt.Println(message...)
}