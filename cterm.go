package cterm


import (
	"bufio"
	"fmt"
	"syscall"
	"unsafe"
)


// Create struct to capture colors and reset byte groups
type colorEscapeCodes struct {
	Black, Red, Green, Yellow, Blue, Magenta, Cyan, White string

	Reset string
}

// Write escape sequence to style text
var Colors = colorEscapeCodes {
	Black: 		"\033[30m",
	Red: 		"\033[31m",
	Green: 		"\033[32m",
	Yellow: 	"\033[33m",
	Blue: 		"\033[34m",
	Magenta: 	"\033[35m",
	Cyan: 		"\033[36m",
	White: 		"\033[37m",
	Reset: 		"\033[0m",
}



// window size class (expected from ioctl)
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


// Sets terminal into pseudo raw-mode. This disables line buffering
func Raw() func() {
	// Get current terminal settings
	var termios syscall.Termios
	_, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TCGETS),
		uintptr(unsafe.Pointer(&termios)),
	)

	// Error handling
	if err != 0 {
		panic(err)
	}

	// Save terminal settings for restore
	oldTermios := termios

	// Change terminal settings
	termios.Lflag &^= syscall.ECHO | syscall.ICANON // Disable echo and canonical mode
	termios.Cc[syscall.VMIN] = 1 					// Minimum read characters
	termios.Cc[syscall.VTIME] = 0					// Infinite timeout

	// Set Raw mode
	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TCSETS),
		uintptr(unsafe.Pointer(&termios)),
	)

	// This is what gets deferred to restore terminal functionality
	return func() {
		// Set terminal settings back to the original values
		syscall.Syscall(
			syscall.SYS_IOCTL,
			uintptr(syscall.Stdin),
			uintptr(syscall.TCSETS),
			uintptr(unsafe.Pointer(&oldTermios)),
		)
	}
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


// Move cursor to specified location
func MoveCursor(buffer *bufio.Writer, x, y int){
	fmt.Fprintf(buffer, "\033[%d;%dH", y+1, x+1)
}
