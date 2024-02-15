package models

type Book struct {
	Title   string
	Author  string
	Country string
}

type BookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
