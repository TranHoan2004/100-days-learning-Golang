package service

import (
	"errors"
	"normal_app/model"
)

type Service struct {
	dictionary *model.Dictionary
}

func NewService() *Service {
	return &Service{
		dictionary: model.NewDictionary(),
	}
}

func (s *Service) AddWord(eng string, vie string) error {
	if err := s.HasWord(eng); err {
		return errors.New("word already exists")
	}
	return s.dictionary.AddWord(
		eng,
		model.NewWord(vie),
	)
}

func (s *Service) UpdateWord(keyword string, content string) string {
	if err := s.HasWord(keyword); !err {
		return "this word is not existing"
	}
	s.dictionary.UpdateWord(
		keyword,
		model.NewWord(content),
	)
	return "update word successfully"
}

func (s *Service) DeleteWord(keyword string) string {
	if err := s.HasWord(keyword); !err {
		return "this word is not existing"
	}
	s.dictionary.DeleteWord(keyword)
	return "delete successfully"
}

func (s *Service) HasWord(w string) bool {
	return s.dictionary.HasWord(w)
}

func (s *Service) GetWords() map[string]string {
	result := make(map[string]string)
	for eng, vie := range s.dictionary.Words {
		result[eng] = vie.GetName()
	}
	return result
}
