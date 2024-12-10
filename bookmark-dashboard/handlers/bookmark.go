package handlers

import (
    "bookmark-dashboard/models"
    "net/http"
    "strconv"
    "html/template"
)

func AddBookmarkFormHandler(w http.ResponseWriter, r *http.Request) {
    categoryID := r.URL.Query().Get("category_id")
    tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/form.html"))
    tmpl.Execute(w, map[string]interface{}{
        "Title":  "Add Bookmark",
        "Action": "/add-bookmark",
        "Fields": []map[string]string{
            {"Name": "category_id", "Label": "Category ID", "Value": categoryID},
            {"Name": "title", "Label": "Bookmark Title", "Value": ""},
            {"Name": "url", "Label": "Bookmark URL", "Value": ""},
        },
    })
}

func AddBookmarkHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        title := r.FormValue("title")
        url := r.FormValue("url")
        categoryID, _ := strconv.Atoi(r.FormValue("category_id"))

        data, _ := loadData()
        for i, category := range data.Categories {
            if category.ID == categoryID {
                newID := len(category.Bookmarks) + 1
                category.Bookmarks = append(category.Bookmarks, models.Bookmark{
                    ID:    newID,
                    Title: title,
                    URL:   url,
                })
                data.Categories[i] = category
                saveData(data)
                break
            }
        }
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}
