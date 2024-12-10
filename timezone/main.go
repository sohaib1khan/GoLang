package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// Structs for Time Zone and Template Data
type ComparisonResult struct {
	ESTTime string
	UTCTime string
}

type PageData struct {
	CurrentEST  string
	CurrentUTC  string
	Comparison  *ComparisonResult
	Error       string
}

// Function to Open Default Browser
func openBrowser(url string) {
	switch runtime.GOOS {
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	case "linux":
		exec.Command("xdg-open", url).Start()
	default:
		fmt.Printf("Cannot open browser for %s\n", runtime.GOOS)
	}
}

// Function to Get Current Time in EST and UTC
func getCurrentTimes() (string, string) {
	utcNow := time.Now().UTC()
	estNow := utcNow.Add(-5 * time.Hour) // Adjust UTC for EST (-5 hours)

	return estNow.Format("03:04 PM"), utcNow.Format("03:04 PM") // 12-hour format
}

// HTTP Handler Function
func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Fetch current times
	currentEST, currentUTC := getCurrentTimes()

	// Variables to hold comparison results
	var comparison *ComparisonResult
	var errorMsg string

	if r.Method == http.MethodPost {
		// Handle form input
		estTimeStr := r.FormValue("estTime")

		// Normalize input: convert to uppercase
		normalizedESTTimeStr := strings.ToUpper(estTimeStr)

		parsedESTTime, err := time.Parse("03:04 PM", normalizedESTTimeStr)
		if err != nil {
			errorMsg = "Invalid EST time format. Use 12-hour format (e.g., 02:30 PM)."
		} else {
			utcTime := parsedESTTime.Add(5 * time.Hour) // Convert EST to UTC
			comparison = &ComparisonResult{
				ESTTime: parsedESTTime.Format("03:04 PM"),
				UTCTime: utcTime.Format("03:04 PM"),
			}
		}
	}

	// Prepare data for template
	data := PageData{
		CurrentEST:  currentEST,
		CurrentUTC:  currentUTC,
		Comparison:  comparison,
		Error:       errorMsg,
	}

	// Render the template
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, data)
}

// Main Function
func main() {
	url := "http://localhost:8080"
	go openBrowser(url) // Open the browser

	http.HandleFunc("/", handleIndex)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Printf("Starting server at %s\n", url)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
