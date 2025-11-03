package main

import (
	"fmt"
	"os"
	"github.com/Nareleth/cterm"
)

func main() {
	// Disable line buffering
	cleanup := cterm.Raw()

	// Run the cleanup actions to restore echoing and line buffering when program finishes
	defer cleanup()

	fmt.Println("Press any key (q to quit):")

	// Initialize a 1 size slice for capturing key input
	key := make([]byte, 1)

	// Infinite loop
	for {
		// Read input from stdin
		os.Stdin.Read(key)

		// Echo back what key was pressed
		fmt.Printf("You pressed: %c\n", key[0])

		// Simple break condition to end program
		if key[0] == 'q' {
			break
		}
	}
}