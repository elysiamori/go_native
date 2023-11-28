package models

// 0. define struct data with respinse json from api
type Content struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Date  string `json:"date"`
}

// get all data
type Contents struct {
	Contents []Content `json:"contents"`
}

// get data by id
// type Contents struct {
// 	Contents Content `json:"content"`
// }
