package main

import (
	"fmt"
	"log"
	"time"

	"github.com/atotto/clipboard"
	"github.com/micmonay/keybd_event"
	"github.com/micmonay/simhotkey"
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
	// Initialize the keyboard event
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		log.Fatalf("Failed to initialize keyboard event: %v\n", err)
	}

	// Simulate typing character by character
	for _, char := range text {
		kb.SetKeys(int(char))
		kb.Launching()
		time.Sleep(50 * time.Millisecond) // Small delay between key presses
	}
}

func main() {
	// Create a new hotkey listener
	listener := simhotkey.NewHotKeyListener()

	// Register hotkeys
	err := listener.RegisterHotKey("ctrl+c", func() {
		fmt.Println("Hotkey Ctrl+C triggered.")
		copyClipboard()
	})
	if err != nil {
		log.Fatalf("Failed to register hotkey Ctrl+C: %v\n", err)
	}

	err = listener.RegisterHotKey("ctrl+1", func() {
		fmt.Println("Hotkey Ctrl+1 triggered.")
		typeSharedDriveURL()
	})
	if err != nil {
		log.Fatalf("Failed to register hotkey Ctrl+1: %v\n", err)
	}

	err = listener.RegisterHotKey("ctrl+2", func() {
		fmt.Println("Hotkey Ctrl+2 triggered.")
		typeDashboardURL()
	})
	if err != nil {
		log.Fatalf("Failed to register hotkey Ctrl+2: %v\n", err)
	}

	fmt.Println("App is running... Press Ctrl+C to stop.")

	// Run the hotkey listener
	listener.Listen()
}
