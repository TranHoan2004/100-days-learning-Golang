package repository

import "Try_to_create_clean_architecture_application/domain/entities"

type BookShelfRepository interface {
	FindAll() ([]entities.BookShelf, error)
	FindById(ID string) (entities.BookShelf, error)
	Save(entity entities.BookShelf) error
	Delete(ID string) error
}
