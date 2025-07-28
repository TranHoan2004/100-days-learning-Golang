package repository

import "Try_to_create_clean_architecture_application/domain/entities"

type BookRepository interface {
	FindAll() ([]entities.Book, error)
	FindById(ID string) (entities.Book, error)
	Save(book entities.Book) error
	Delete(ID string) error
}
