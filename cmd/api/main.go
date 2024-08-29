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

	http.HandleFunc("/hello", greet)
	http.HandleFunc("/books", bookHandlers.GetBooks)
	http.ListenAndServe(":8080", nil)
}
