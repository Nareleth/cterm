# Terminal Canvas (cterm)
## Overview
Take control of your terminal and elevate it to a proper engine.  
Set your terminal to Raw Mode, print colored text, get your terminal size, and control the cursor.

## Installation
1. Initialize your project (if you haven't done so).  
```go mod init your/project/name```  
1. Import this package in your project: ```github.com/Nareleth/cterm```
1. Get this package  
```go mod tidy```  

## Examples
For quick start references, see the ```examples``` directory for usage.  
Set Terminal in Raw Mode:  
```
// Set terminal in Raw Mode
cleanup := cterm.Raw()

// Restore terminal settings when finished
defer cleanup()
```


## API Reference
### Variables
#### var Colors
```
var Colors  
 ```  
<details>
<summary>Example</summary>
```
fmt.Println(cterm.Colors(Red) + "Red" + cterm.Colors.Reset)
```
</details>


### Functions
#### func Clear
```
func Clear(screen *bufio.Writer)
 ```  
 Clear uses ANSI escape codes to clear screen buffer.  
<details>
<summary>Example</summary>
```
cterm.Clear(screen)
```
</details>


#### func GetSize
```
func GetSize() (int, int, error)  
```  
Return the visible dimensions of your terminal.  
<details>
<summary>Example</summary>
```
width, height, _ := cterm.GetSize()
fmt.Printf("Terminal Dimensions: (%d, %d)\n", width, height)
```
</details>


#### func HideCursor
```
func HideCursor(screen *bufio.Writer)  
```  
HideCursor hides the display of your cursor
<details>
<summary>Example</summary>
```
cterm.HideCursor(screen)
```
</details>


#### func MoveCursor
```
func MoveCursor(buffer *bufio.Writer, x, y int)
```  
MoveCursor sets the position of your cursor to a specified location, usually for printing.
<details>
<summary>Example</summary>
```
// Move cursor to top of screen
cterm.MoveCursor(screen, 1, 1)
```
</details>


#### func Raw
```
func Raw() func()  
```  
Raw sets the terminal to raw mode, and restores to previous state when finished.  
<details>
<summary>Example</summary>
```
// Set terminal in Raw Mode
cleanup := cterm.Raw()
// Restore terminal settings when finished
defer cleanup()
```
</details>


#### func ShowCursor
```
func  ShowCursor(screen *bufio.Writer)  
 ```  
<details>
ShowCursor echos your cursor position if it was hidden.   
<summary>Example</summary>
```
cterm.ShowCursor(screen)
```
</details>

