package main

import (
	"crud-books/internal/service"
	"crud-books/internal/web"

	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	db, err := sql.Open("sqlite3", "./infra/db.sql")

	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	defer db.Close()

	bookService := service.NewBookService(db)

	bookHandlers := web.NewBookHandlers(bookService)

	router := http.NewServeMux()

	router.HandleFunc("GET /hello", greet)
	router.HandleFunc("GET /books", bookHandlers.GetBooks)
	router.HandleFunc("GET /books/{id}", bookHandlers.GetBookByID)
	router.HandleFunc("POST /books", bookHandlers.CreateBook)
	router.HandleFunc("PUT /books/{id}", bookHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandlers.DeleteBook)

	http.ListenAndServe(":8080", router)
}
