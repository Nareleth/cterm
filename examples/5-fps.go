package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/Nareleth/cterm"
)

// Create a screen buffer
var screen = bufio.NewWriter(os.Stdout)

func main() {
	// Starts new game clock at 30fps
	gameClock := cterm.NewClock(30) 

	// Disable line buffering
	cleanup := cterm.Raw()

	// Run the cleanup actions to restore echoing and line buffering when program finishes
	defer cleanup()


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
		fmt.Fprintf(screen, "FPS: %d\nPress q to exit\n", gameClock.GetFPS())

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