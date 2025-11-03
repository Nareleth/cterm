package main


import (
	"fmt"
	"strings"
	"github.com/Nareleth/cterm"
)


func main() {
	// Initialize variables

	// Unicode block character
	block := "\u2588"
	// Slice of colors
	colors 		:= []string{"Black", "Red", "Yellow", "Blue", "Magenta", "Cyan", "White"}
	// Slice of functions to display color
	colorCodes 	:= []string{cterm.Colors.Black, cterm.Colors.Red, cterm.Colors.Yellow, cterm.Colors.Blue, cterm.Colors.Magenta, cterm.Colors.Cyan, cterm.Colors.White}

	// Loop iterate through all the colors. Printing the name and blocks all while colored, then reset to default color
	for i := 0; i < len(colors); i++ {
		fmt.Printf("%s%-7s %s%s\n", 
		colorCodes[i], colors[i], 
		strings.Repeat(block, 10),
		cterm.Colors.Reset)
	}
}