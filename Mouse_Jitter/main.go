package main

import (
	"fmt"
	"time"
	"syscall"
	"unsafe"
)

// Import Windows API functions
var (
	user32           = syscall.NewLazyDLL("user32.dll")
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	getForegroundWindow = user32.NewProc("GetForegroundWindow")
	findWindow        = user32.NewProc("FindWindowW")
	setCursorPos      = user32.NewProc("SetCursorPos")
)

// isRDPRunning checks if the RDP window (mstsc.exe) is running by its window title.
func isRDPRunning() bool {
	// Find the RDP window by its title (can vary based on locale or session state)
	hWnd, _, _ := findWindow.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Remote Desktop Connection"))), 0)
	return hWnd != 0
}

// preventLock simulates mouse movement to keep the session active.
func preventLock() {
	// Move the mouse slightly to prevent screen lock
	fmt.Println("Moving mouse to prevent lock...")
	for {
		// Simulate small mouse movements
		setCursorPos.Call(100, 100) // Move to position 100, 100
		time.Sleep(500 * time.Millisecond)
		setCursorPos.Call(200, 200) // Move to position 200, 200
		time.Sleep(30 * time.Second) // Wait 30 seconds before moving again
	}
}

func main() {
	fmt.Println("Starting RDP Lock Preventer...")
	for {
		if isRDPRunning() {
			fmt.Println("RDP detected! Preventing lock...")
			preventLock()
		} else {
			fmt.Println("RDP not running.")
			time.Sleep(10 * time.Second) // Check again after 10 seconds
		}
	}
}


// 
