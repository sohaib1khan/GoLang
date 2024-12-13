package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2" // Fiber is a web framework for building web applications
)

// Template struct is used to manage HTML templates
type Template struct {
	templates *template.Template
}

// Render renders an HTML template to the browser
func (t *Template) Render(w io.Writer, name string, data interface{}, c ...string) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Load loads the templates from the views folder
func (t *Template) Load() error {
	// Parse all HTML files in the views directory
	t.templates = template.Must(template.ParseGlob("./views/*.html"))
	return nil
}

// Bookmark represents a single bookmark entry
// Each bookmark has a title, a URL, and a category
type Bookmark struct {
	Title    string `json:"title"`    // Title of the bookmark
	URL      string `json:"url"`      // URL of the bookmark
	Category string `json:"category"` // Category to which the bookmark belongs
}

// Data struct holds the categories and bookmarks
// This is what we save and load from the JSON file
type Data struct {
	Categories []string   `json:"categories"` // List of categories
	Bookmarks  []Bookmark `json:"bookmarks"`  // List of bookmarks
}

// Path to the JSON file where data is stored
const DataFile = "data.json"

// Global variables
var data Data        // Holds all the data (categories + bookmarks)
var mu sync.Mutex    // Mutex ensures thread-safe access to the data

// LoadData reads data from the JSON file into the global `data` variable
func LoadData() error {
	file, err := os.Open(DataFile) // Open the data file
	if os.IsNotExist(err) {
		// If the file doesn't exist, initialize it with default values
		data = Data{
			Categories: []string{"Development", "Entertainment"}, // Default categories
			Bookmarks: []Bookmark{
				{Title: "GitHub", URL: "https://github.com", Category: "Development"},
				{Title: "Plex", URL: "https://plex.tv", Category: "Entertainment"},
			},
		}
		return SaveData() // Save the defaults to a new file
	}
	if err != nil {
		return err // Return any other errors
	}
	defer file.Close()

	// Decode the JSON file into the `data` variable
	return json.NewDecoder(file).Decode(&data)
}

// SaveData writes the current state of `data` to the JSON file
func SaveData() error {
	file, err := os.Create(DataFile) // Create (or overwrite) the data file
	if err != nil {
		return err // Return if there's an error creating the file
	}
	defer file.Close()

	// Encode `data` as JSON and write it to the file
	return json.NewEncoder(file).Encode(data)
}

func main() {
	// Load data from the JSON file when the application starts
	if err := LoadData(); err != nil {
		log.Fatalf("Failed to load data: %v", err) // Log an error and stop the program if loading fails
	}

	// Initialize the template renderer
	tmpl := &Template{}

	// Initialize the Fiber web application
	app := fiber.New(fiber.Config{
		Views: tmpl, // Set the template renderer for HTML views
	})

	// Home route - Renders the main HTML page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index.html", nil) // Render the "index.html" file
	})

	// API route to get all data (categories and bookmarks)
	app.Get("/api/data", func(c *fiber.Ctx) error {
		mu.Lock()             // Lock the data for thread-safe access
		defer mu.Unlock()     // Unlock after the function finishes
		return c.JSON(data)   // Return the data as JSON
	})

	// API route to update all data (categories and bookmarks)
	app.Post("/api/data", func(c *fiber.Ctx) error {
		var newData Data
		// Parse the JSON request body into the `newData` variable
		if err := c.BodyParser(&newData); err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid data format")
		}

		mu.Lock()            // Lock the data for thread-safe modification
		data = newData       // Replace the current data with the new data
		mu.Unlock()          // Unlock after the modification

		// Save the updated data to the JSON file
		if err := SaveData(); err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Failed to save data")
		}

		return c.SendStatus(http.StatusOK) // Return a 200 OK status
	})

	// Serve static files (e.g., CSS and JavaScript)
	app.Static("/static", "./static")

	// Start the server on port 3000
	log.Println("Server running on http://localhost:3002")
	log.Fatal(app.Listen(":3002")) // Log any errors that occur while starting the server
}
