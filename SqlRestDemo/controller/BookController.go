package controller

import (
	"SqlRestDemo/dao"
	"encoding/json"
	"fmt"
	"net/http"
)

//BookController is to call functions in BookDao
type BookController struct {
}

var bookDao = new(dao.BookDao)

//GetAllBooks is to query all the books from database
func (*BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COntent-Type", "application/json")
	var books = bookDao.GetAllBooks()
	fmt.Println(books)
	json.NewEncoder(w).Encode(books)
}
