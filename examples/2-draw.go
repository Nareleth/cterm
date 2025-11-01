package main

import(
	"bufio"
	"os"
	"github.com/Nareleth/cterm"
)


var screen = bufio.NewWriter(os.Stdout)


func main() {
	cterm.Clear(screen)
	cterm.HideCursor(screen)

	
	cterm.ShowCursor(screen)
	screen.Flush()
}