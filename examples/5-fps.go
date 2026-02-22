package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/Nareleth/cterm"
)

// Create a screen buffer
var screen = bufio.NewWriter(os.Stdout)

/*
// Create a Cursor object for Delta time testing
type Cursor struct {
	x 		int
	y 		int
	sprite 	rune
	speed 	int
}


// Create a reusable movement function for the cursors
func (p *Cursor) Move() {
	return p.y += p.speed
}
*/


// Render function to keep main cleaner
func render() {
	col := 0

	// Initialize a screen grid
	for i := 0; i < col; i++ {
		fmt.Fprint(screen)
	}


	fmt.Fprint(screen, "#")
}



func main() {
	// Starts new game clock at 30fps
	gameClock := cterm.NewClock(30) 

	// Disable line buffering
	cleanup := cterm.Raw()

	// Run the cleanup actions to restore echoing and line buffering when program finishes
	defer cleanup()

/*
	// Declare cursor objects to test delta speed
	cursorTestBase := &Cursor {
		x: 		0,
		y: 		2,
		sprite: '\u2588',
		speed:	5,
	}

	cursorTestDelta := &Cursor {
		x: 		10,
		y: 		2,
		sprite: '\u2588',
		speed:	5,
	}
*/

	// Channel to read key press
	key := make(chan byte)

	// Goroutine to read keypress
	go func() {
		keyPress := make([]byte, 1)
		for {
			os.Stdin.Read(keyPress)
			key <- keyPress[0]
		}
	}()


	// Start animation loop
	for {
		// Clear the screen
		cterm.Clear(screen)

		// Starts frame calculating
		gameClock.FrameStart()

		// Render FPS reader
		cterm.MoveCursor(screen, 0, 0)
		fmt.Fprintf(screen, "FPS: %d\n", gameClock.GetFPS())

		// Render quit text
		cterm.MoveCursor(screen, 0, 1)
		fmt.Fprintf(screen, "Press q to exit\n")


		// Render/Drawing layer
		render()
		

		// Clears the buffer data and writes to STDOUT
		screen.Flush()

		// Key press event listener
		select {
		case keyPress := <-key:
			if keyPress == 'q' {
				return
			}
		default:
			// Nothing
		}

		// Ends frame calculating
		gameClock.FrameEnd()
	}

}