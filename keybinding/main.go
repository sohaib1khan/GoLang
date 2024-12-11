package main

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
)

func copyClipboard() {
	// Read text from the clipboard
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Printf("Failed to copy from clipboard: %v\n", err)
		return
	}
	// Simulate typing the text
	typeText(text)
	fmt.Println("Copied text from clipboard and typed it.")
}

func typeSharedDriveURL() {
	// Shared drive URL
	sharedDriveURL := `http://example.com/dashboard`
	typeText(sharedDriveURL)
	fmt.Println("Typed shared drive URL.")
}

func typeDashboardURL() {
	// Dashboard URL
	dashboardURL := "http://example.com/dashboard"
	typeText(dashboardURL)
	fmt.Println("Typed dashboard URL.")
}

func typeText(text string) {
	// Type text character by character
	for _, char := range text {
		robotgo.TypeStr(string(char))
		time.Sleep(50 * time.Millisecond) // Delay to mimic natural typing
	}
}

func main() {
	// Register global hotkeys
	fmt.Println("Registering hotkeys...")

	// Ctrl+C to copy clipboard and type
	robotgo.EventHook(robotgo.KeyDown, []string{"ctrl", "c"}, func(e robotgo.Event) {
		fmt.Println("Hotkey Ctrl+C triggered.")
		copyClipboard()
	})

	// Ctrl+1 to type shared drive URL
	robotgo.EventHook(robotgo.KeyDown, []string{"ctrl", "1"}, func(e robotgo.Event) {
		fmt.Println("Hotkey Ctrl+1 triggered.")
		typeSharedDriveURL()
	})

	// Ctrl+2 to type dashboard URL
	robotgo.EventHook(robotgo.KeyDown, []string{"ctrl", "2"}, func(e robotgo.Event) {
		fmt.Println("Hotkey Ctrl+2 triggered.")
		typeDashboardURL()
	})

	// Keep the program running to listen for hotkeys
	fmt.Println("Listening for hotkeys... Press Ctrl+Alt+Q to exit.")
	robotgo.EventHook(robotgo.KeyDown, []string{"ctrl", "alt", "q"}, func(e robotgo.Event) {
		fmt.Println("Exiting program.")
		robotgo.StopEvent()
	})
	robotgo.EventStart()
}
