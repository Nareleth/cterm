package main


import (
	"fmt"
	"github.com/Nareleth/cterm"
)


// Get the dimensions of the terminal
func main() {
	// Get width and height from the GetSize function
	width, height, _ := cterm.GetSize()

	// Print the Terminal Dimensions in (X, Y) plot form
	fmt.Printf("Terminal Dimensions: (%d, %d)\n", width, height)
}