package main

import (
	"fmt"
	"strings"
	"time"
	"syscall"
	"unsafe"
)

// Import Windows API functions
var (
	user32        = syscall.NewLazyDLL("user32.dll")
	kernel32      = syscall.NewLazyDLL("kernel32.dll")
	findWindow    = user32.NewProc("FindWindowW")
	setCursorPos  = user32.NewProc("SetCursorPos")
	enumProcesses = kernel32.NewProc("K32EnumProcesses")
)

// isProcessRunning checks if a specific process is running
func isProcessRunning(targets []string) bool {
	for _, title := range targets {
		hWnd, _, _ := findWindow.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0)
		if hWnd != 0 {
			fmt.Printf("Detected process: %s\n", title)
			return true
		}
	}
	return false
}

// preventLock simulates mouse movement to keep the session active
func preventLock() {
	fmt.Println("Moving mouse to prevent lock...")
	for {
		setCursorPos.Call(100, 100) // Move to position 100, 100
		time.Sleep(500 * time.Millisecond)
		setCursorPos.Call(200, 200) // Move to position 200, 200
		time.Sleep(30 * time.Second) // Wait 30 seconds before moving again
	}
}

func main() {
	// List of process titles to monitor
	targetProcesses := []string{"mstsc", "iexplore", "msedge", "sophia", "ia"}

	fmt.Println("Starting process monitor...")

	for {
		if isProcessRunning(targetProcesses) {
			fmt.Println("Detected target process. Preventing lock...")
			preventLock()
		} else {
			fmt.Println("No target process running.")
			time.Sleep(10 * time.Second) // Check again after 10 seconds
		}
	}
}
