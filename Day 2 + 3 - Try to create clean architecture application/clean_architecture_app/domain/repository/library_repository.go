package repository

import "Try_to_create_clean_architecture_application/domain/entities"

type LibraryRepository interface {
	FindAll() ([]entities.Library, error)
	FindById(ID string) (entities.Library, error)
	Save(entity entities.Library) error
	Delete(ID string) error
}
