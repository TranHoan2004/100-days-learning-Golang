package business

import (
	"Try_to_create_clean_architecture_application/domain/entities"
	"Try_to_create_clean_architecture_application/domain/repository"
)

type BookService struct {
	repo repository.BookRepository
}

func CreateNewBook(book entities.Book) {
	
}