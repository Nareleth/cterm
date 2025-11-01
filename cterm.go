package cterm

import (
	"bufio"
	"fmt"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row		uint16
	Col		uint16
	Xpixel	uint16
	Ypixel 	uint16
}


// Get the size of the terminal window
func GetSize() (int, int, error){
	// Init and create pointer to winsize
	ws := &winsize{}

	// SYSCALL IOCTL - GET THE WIN SIZE
	returnCode, _, errorCode := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	// Error handling
	if int(returnCode) == -1 {
		return 0, 0, errorCode
	}

	// Return columns and rows (width, height)
	return int(ws.Col), int(ws.Row), nil
}


// Clear the screen using raw terminal escape code
func Clear(screen *bufio.Writer){
	fmt.Fprint(screen, "\033c")
}


// Hide the Cursor
func HideCursor(screen *bufio.Writer){
	fmt.Fprint(screen, "\033[?25l")
}


// Show the Cursor
func ShowCursor(screen *bufio.Writer){
	fmt.Fprint(screen, "\033[?25h")
}


/*
// Move cursor to specified location
func MoveCursor(buffer io.Writer, x, y int){
	fmt.Fprintf(buffer, "\033[%d;%dH", y+1, x+1)
}