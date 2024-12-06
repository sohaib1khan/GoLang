package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/shirou/gopsutil/process"
)

func isRDPRunning() bool {
	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Error fetching processes:", err)
		return false
	}

	for _, proc := range processes {
		name, err := proc.Name()
		if err == nil && name == "mstsc.exe" {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Starting RDP Monitor...")

	for {
		if isRDPRunning() {
			fmt.Println("RDP is running. Preventing lock...")
			// Move the mouse slightly to prevent lock
			x, y := robotgo.GetMousePos()
			robotgo.MoveMouse(x+1, y+1)
			time.Sleep(100 * time.Millisecond)
			robotgo.MoveMouse(x, y)
			time.Sleep(30 * time.Second) // Wait before moving again
		} else {
			fmt.Println("RDP is not running.")
			time.Sleep(10 * time.Second) // Check again after 10 seconds
		}
	}
}
