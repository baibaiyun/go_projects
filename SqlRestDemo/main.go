package main

import (
	"SqlRestDemo/controller"
	"SqlRestDemo/dao"
	"SqlRestDemo/utils"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/mattn/go-adodb"
)

func main() {
	var connectDb = new(dao.ConnectDB)
	utils.DB = connectDb.InitDB()

	var bookController = new(controller.BookController)

	//Init router
	r := mux.NewRouter()
	//Create router handler - Endpoint
	r.HandleFunc("/api/books", bookController.GetAllBooks).Methods("GET")
	http.ListenAndServe(":3000", r)

	defer utils.DB.Close()
}
