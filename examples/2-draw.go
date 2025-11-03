package main

import(
	"bufio"
	"fmt"
	"os"
	"github.com/Nareleth/cterm"
)


// Create a buffer for the console
var screen = bufio.NewWriter(os.Stdout)



func main() {
	// Initialize vars
	// Get width and height from the GetSize function
	width, height, _ := cterm.GetSize()

	// Initialize a string value for the box we will draw in this example
	border 	:= "#"
	body	:= "-"


	// Clear the screen so we can only see what we are drawing
	cterm.Clear(screen)

	// Draw to our screen buffer using the X axis (columns)
	for col := 0; col < width; col++ {
		// Move the cursor to a specified location using the current position of the loop
		cterm.MoveCursor(screen, col, 0)

		// Print the string of our border to the screen buffer
		fmt.Fprint(screen, border)
		
		// Repeat to Draw the bottom border
		cterm.MoveCursor(screen, col, height - 2)
		fmt.Fprint(screen, border)
	}

	// Draw to our screen buffer using the Y axis (rows)
	for row := 0; row < height - 2; row++ {
		// Move the cursor to a specified location using the current position of the loop
		cterm.MoveCursor(screen, 0, row)

		// Print the string of our border to the screen buffer
		fmt.Fprint(screen, border)
		
		// Repeat to Draw the bottom border
		cterm.MoveCursor(screen, width - 1, row)
		fmt.Fprint(screen, border)
	}

	// Grid based loop - iterate through the columns (X), then the Rows (Y)
	// We want to draw inside the box, so we change the values of col and row
	for col := 1; col < width - 1; col++ {
		for row := 1; row < height - 2; row ++ {
			// We print to the current position in the nested loop
			cterm.MoveCursor(screen, col, row)
			fmt.Fprint(screen, body)
		}
	}
	
	// Move cursor out of our buffer
	cterm.MoveCursor(screen, 0, height + 1)

	// Clears the buffer data  and writes to STDOUT
	screen.Flush()
}