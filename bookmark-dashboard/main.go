package main

import (
	"bookmark-dashboard/handlers"
	"log"
	"net/http"
)

func main() {
	// Initialize JSON storage
	handlers.InitStorage("data/data.json")

	// Routes
	http.HandleFunc("/", handlers.DashboardHandler)
	http.HandleFunc("/add-category-form", handlers.AddCategoryFormHandler)
	http.HandleFunc("/add-category", handlers.AddCategoryHandler)
	http.HandleFunc("/add-bookmark-form", handlers.AddBookmarkFormHandler)
	http.HandleFunc("/add-bookmark", handlers.AddBookmarkHandler)


	// Static assets (CSS, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server is running on http://localhost:8181")
	log.Fatal(http.ListenAndServe(":8181", nil))
}
