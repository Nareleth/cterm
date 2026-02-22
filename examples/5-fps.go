package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"github.com/Nareleth/cterm"
)

// Create a screen buffer
var screen = bufio.NewWriter(os.Stdout)


// Create a Cursor object for Delta time testing
type Cursor struct {
	x 		int
	y 		int
	sprite 	rune
	speed 	int
}


// Create a reusable movement function for the cursors
func (c *Cursor) Move(startY, maxY int, deltaTime float64) {

	// Check if delta  time is used
	if deltaTime > 0 {
		c.y += int(float64(c.speed) * deltaTime)
	} else {
		c.y += c.speed
	}

	if c.y > maxY {
		c.y = startY
	}
}


// Create a draw function for the cursors
func (c *Cursor) Draw() {
	cterm.MoveCursor(screen, c.x, c.y)
	fmt.Fprintf(screen, "%c", c.sprite)
}



func main() {
	// Configure me!
	SetFPS := 30

	// Starts new game clock at 30fps
	gameClock := cterm.NewClock(SetFPS) 

	// Disable line buffering
	cleanup := cterm.Raw()

	// Run the cleanup actions to restore normal mode
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cleanup()
		os.Exit(0)
	}()

	defer cleanup()


	// Movement track
	startY 	:= 3
	maxY 	:= 10

	cursorRacerAlpha := &Cursor {
		x: 		0,
		y: 		startY,
		sprite: '@',
		speed:	1,
	}

	cursorRacerDelta := &Cursor {
		x: 		10,
		y: 		startY,
		sprite: '@',
		speed:	SetFPS,
	}


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

		// Get Delta Time
		dt := gameClock.GetDeltaTime()

		// Render FPS reader
		cterm.MoveCursor(screen, 0, 0)
		fmt.Fprintf(screen, "FPS: %d\n", gameClock.GetFPS())

		// Render quit text
		cterm.MoveCursor(screen, 0, 1)
		fmt.Fprintf(screen, "Press q to exit\n")


		// Render and draw cursors
		cursorRacerAlpha.Move(startY, maxY, 0)
		cursorRacerAlpha.Draw()

		cursorRacerDelta.Move(startY, maxY, dt)
		cursorRacerDelta.Draw()

		// move cursor to bottom
		cterm.MoveCursor(screen, 0, maxY+1)

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