package models

type Data struct {
    Categories []Category `json:"categories"`
}

type Category struct {
    ID        int        `json:"id"`
    Name      string     `json:"name"`
    Bookmarks []Bookmark `json:"bookmarks"`
}
