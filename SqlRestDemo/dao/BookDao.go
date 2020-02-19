package dao

import (
	"SqlRestDemo/model"
	"SqlRestDemo/utils"
	"fmt"
	"log"
	"strconv"

	//Only need init package
	_ "github.com/mattn/go-adodb"
)

//BookDao is for Get, Insert, Update, Delete the book
type BookDao struct {
}

//GetAllBooks will query all the books from table - Book
func (*BookDao) GetAllBooks() []model.Book {
	rows, err := utils.DB.Query("select * from Book")
	if err != nil {
		fmt.Println("BookDao.getAllBooks() error")
		log.Println(err)
		return nil
	}

	var books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Isbn, &book.Title, &book.Firstname, &book.Lastname)
		fmt.Println(strconv.Itoa(book.ID), strconv.Itoa(book.Isbn), &book.Title, &book.Firstname, &book.Lastname)
		if err != nil {
			fmt.Println("SelectAllBooks error")
			continue
		}
		books = append(books, book)
	}
	rows.Close()
	return books
}
