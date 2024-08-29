package web

import (
	"crud-books/internal/service"
	"encoding/json"
	"net/http"
)

type BookHandlers struct {
	service *service.BookService
}

// NewBookHandlers cria uma nova instância de BookHandlers.
func NewBookHandlers(service *service.BookService) *BookHandlers {
	return &BookHandlers{service: service}
}

func (h *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetBooks()
	if err != nil {
		http.Error(w, "failed to get books", http.StatusInternalServerError)
		return
	}

	if len(books) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
