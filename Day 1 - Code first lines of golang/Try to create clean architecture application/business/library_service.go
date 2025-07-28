package business

import (
	"Try_to_create_clean_architecture_application/domain/entities"
	"Try_to_create_clean_architecture_application/domain/repository"
)

type LibraryService struct {
	repo repository.LibraryRepository
}

func NewLibraryService(r repository.LibraryRepository) *LibraryService {
	return &LibraryService{repo: r}
}

func (s *LibraryService) GetAll() ([]entities.Library, error) {
	return s.repo.FindAll();
}

func (s *LibraryService) Create (lib entities.Library) error {
	return s.repo.Save(lib);
}
