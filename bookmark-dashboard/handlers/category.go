package handlers

import (
    "bookmark-dashboard/models"
    "html/template"
    "net/http"
)

func AddCategoryFormHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/form.html"))
    tmpl.Execute(w, map[string]interface{}{
        "Title":  "Add Category",
        "Action": "/add-category",
        "Fields": []map[string]string{
            {"Name": "name", "Label": "Category Name", "Value": ""},
        },
    })
}

func AddCategoryHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        name := r.FormValue("name")
        data, _ := loadData()
        newID := len(data.Categories) + 1
        data.Categories = append(data.Categories, models.Category{
            ID:        newID,
            Name:      name,
            Bookmarks: []models.Bookmark{},
        })
        saveData(data)
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}
