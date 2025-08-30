package logger

import (
	"math/rand"

	"github.com/fatih/color"
)

var Verbose bool = false

var colors = []color.Attribute{
	color.FgHiRed,
	color.FgHiGreen,
	color.FgHiYellow,
	color.FgHiBlue,
	color.FgHiMagenta,
	color.FgHiCyan,
}

func Log(info any) {
	if Verbose {
		blue := color.RGB(52, 64, 235)
		blue.Printf("[INFO] %v\n", info)
	}
}

func Warn(info any) {
	if Verbose {
		yellow := color.RGB(235, 205, 52)
		yellow.Printf("[WARN] %v\n", info)
	}
}
func Error(info any) {
	red := color.RGB(235, 67, 52)
	red.Printf("[ERROR] %v\n", info)
}
func Cute(info any) {
	if Verbose {
		cute := color.RGB(235, 52, 235)
		cute.Printf("[INFO] %v\n", info)
	}
}

func RandomColor(info string) {
	c := color.New(colors[rand.Intn(len(colors))])
	c.Println(info)
}
