package models

// 1. Define struct for JSON file
type Book struct {
	ID    string `json:"id"`
	Uuid  string `json:"uuid"`
	Title string `json:"title"`
	Desc  string `json:"description"`
	Date  string `json:"date"`
}

type Library struct {
	Library []Book `json:"libraries"`
}
