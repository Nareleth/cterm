package cterm


import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)


// Create struct to capture colors and reset byte groups
type colorEscapeCodes struct {
	Black, Red, Green, Yellow, Blue, Magenta, Cyan, White string

	Reset string
}


// Create a clock struct for FPS
type Clock struct {
	targetFPS 		int 			// User-defined FPS limit
	frameTime 		time.Duration 	// Time to maintain target FPS
	frameStart 		time.Time 		// Last frame

	lastFrameTime 	time.Time 		// Tracks time for delta time
	deltaTime 		float64 		// Delta time

	fpsTimer 		time.Time 		// Timer to track frames per second
	frameCount		int 			// Counter to track frames per second
	currentFPS 		int 			// Actual frames per second
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
	fd := int(os.Stdin.Fd())
	var termios syscall.Termios

	// Get current settings
	if _, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(fd),
		uintptr(syscall.TCGETS),
		uintptr(unsafe.Pointer(&termios)),
	); errno != 0 {
		fmt.Fprintf(os.Stderr, "Failed to get terminal settings: %v\n", errno)
		return func() {} // Return no-op cleanup
	}

	oldState := termios

	// Disable canonical mode and echo
	termios.Lflag &^= syscall.ECHO | syscall.ICANON
	termios.Cc[syscall.VMIN] = 1
	termios.Cc[syscall.VTIME] = 0

	// Apply settings
	if _, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(fd),
		uintptr(syscall.TCSETS),
		uintptr(unsafe.Pointer(&termios)),
	); errno != 0 {
		fmt.Fprintf(os.Stderr, "Failed to set raw mode: %v\n", errno)
		return func() {} // Return no-op cleanup
	}

	// Return restore function
	return func() {
		syscall.Syscall(
			syscall.SYS_IOCTL,
			uintptr(fd),
			uintptr(syscall.TCSETS),
			uintptr(unsafe.Pointer(&oldState)),
		)
	}
}


// Clear the screen using raw terminal escape code
func Clear(screen *bufio.Writer){
	fmt.Fprint(screen, "\033[2J\033[H")
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


// Clock and FPS


// Create a new game clock for target frames per second
func NewClock(targetFPS int) *Clock {
	currentTime := time.Now()

	return &Clock{
		targetFPS: 		targetFPS,	
		frameTime: 		time.Second / time.Duration(targetFPS),
		frameStart: 	currentTime,
		lastFrameTime: 	currentTime,
		fpsTimer:		currentTime,
	}		
}



// Start a new frame
func (c *Clock) FrameStart() {
	now := time.Now()

	// Calculate Delta Time
	c.deltaTime = now.Sub(c.lastFrameTime).Seconds()
	c.lastFrameTime = now

	c.frameStart = time.Now()
	c.frameCount++

	// Calculate FPS
	if time.Since(c.fpsTimer) >= time.Second {
		// FPS
		c.currentFPS = c.frameCount

		// Reset timer and frames
		c.frameCount = 0
		c.fpsTimer = time.Now()

	}
}


// Ends the frame
func (c *Clock) FrameEnd() {
	// Check elapsed time since the frame start
	elapsedTime := time.Since(c.frameStart)

	// If the frame rendered too fast, wait for target ms
	if elapsedTime < c.frameTime {
		time.Sleep(c.frameTime - elapsedTime)
	}
}


// Returns the value of the current FPS from the game clock
func (c *Clock) GetFPS() int {
	return c.currentFPS
}


// Returns Delta time for consistent measurements
func (c *Clock) GetDeltaTime() float64 {
	return c.deltaTime
}