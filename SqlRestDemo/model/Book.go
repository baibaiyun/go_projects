package model

type Book struct {
	ID        int    `json:"id"`
	Isbn      int    `json:"isbn"`
	Title     string `json:"title"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
