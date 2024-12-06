package main

import (
	"log"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// Loop to periodically send keystrokes
	for {
		// Get the active window (optional, check if it's the target)
		activeWindow := robotgo.GetTitle()
		log.Printf("Active Window: %s", activeWindow)

		// Send a random keystroke (simulate activity)
		robotgo.TypeStr(" ")
		log.Println("Sent keystroke to keep session active")

		// Wait for a few seconds before sending the next keystroke
		time.Sleep(30 * time.Second)
	}
}
