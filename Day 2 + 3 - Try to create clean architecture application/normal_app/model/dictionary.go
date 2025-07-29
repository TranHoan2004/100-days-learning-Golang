package model

import (
	"errors"
	"github.com/google/uuid"
)

type Dictionary struct {
	ID    uuid.UUID
	Words map[string]Word
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		ID:    uuid.New(),
		Words: make(map[string]Word),
	}
}

func NewDictionaryWithArgs(words map[string]Word) *Dictionary {
	return &Dictionary{
		ID:    uuid.New(),
		Words: words,
	}
}

func (d *Dictionary) GetID() uuid.UUID {
	return d.ID
}

func (d *Dictionary) GetWords() map[string]Word {
	return d.Words
}

func (d *Dictionary) SetID() {
	d.ID = uuid.New()
}

func (d *Dictionary) SetWords(words map[string]Word) {
	d.Words = words
}

func (d *Dictionary) AddWord(eng string, vie Word) error {
	if eng == "" && vie.GetName() == "" {
		return errors.New("words must not be empty")
	}
	d.Words[eng] = vie
	return nil
}

func (d *Dictionary) UpdateWord(keyword string, target Word) {
	d.Words[keyword] = target
}

func (d *Dictionary) DeleteWord(keyword string) {
	delete(d.Words, keyword)
}

func (d *Dictionary) HasWord(keyword string) bool {
	_, existing := d.Words[keyword]
	return existing
}
