// Copyright (c) 2019, Rail Inc.
// All rights reserved.
// Description: Go project.
// Author: baibiayun
// Date: 2020/02/15

package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init books var as a slice Book struct
var books []Book

//Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COntent-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COntent-Type", "application/json")
	params := mux.Vars(r) //get params
	//Loop through books and find the ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//Create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COntent-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //create ID for new book
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COntent-Type", "application/json")
	params := mux.Vars(r) //get params
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] //create ID for new book
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	//json.NewEncoder(w).Encode(books)
}

//Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COntent-Type", "application/json")
	params := mux.Vars(r) //get params
	for index, item := range books {
		if item.ID == params["ID"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	//Init router
	r := mux.NewRouter()

	//Mock data - @todo - implement DB
	books = append(books, Book{
		ID: "1", Isbn: "448743", Title: "Book one", Author: &Author{
			Firstname: "John", Lastname: "Doe",
		},
	})
	books = append(books, Book{
		ID: "2", Isbn: "448742", Title: "Book two", Author: &Author{
			Firstname: "John", Lastname: "Doe2",
		},
	})

	//Create router handler - Endpoint
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	http.ListenAndServe(":3000", r)
}
