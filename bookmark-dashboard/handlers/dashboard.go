package handlers

import (
    "html/template"
    "net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
    data, err := loadData()
    if err != nil {
        http.Error(w, "Failed to load data", http.StatusInternalServerError)
        return
    }

    tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/dashboard.html"))
    tmpl.Execute(w, data.Categories) // Pass categories to the template
}
