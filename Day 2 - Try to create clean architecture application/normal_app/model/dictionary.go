package model

import (
	"errors"
	"github.com/google/uuid"
)

type Dictionary struct {
	ID    uuid.UUID
	Words map[Word]Word
}

func NewDictionary(words map[Word]Word) *Dictionary {
	return &Dictionary{
		ID:    uuid.New(),
		Words: words,
	}
}

func (d *Dictionary) GetID() uuid.UUID {
	return d.ID
}

func (d *Dictionary) GetWords() map[Word]Word {
	return d.Words
}

func (d *Dictionary) SetID() {
	d.ID = uuid.New()
}

func (d *Dictionary) SetWords(words map[Word]Word) {
	d.Words = words
}

func (d *Dictionary) AddWord(eng Word, vie Word) error {
	if eng.GetName() == "" && vie.GetName() == "" {
		return errors.New("words must not be empty")
	}
	if status := d.HasWord(eng); status {
		return errors.New("this word is already existing")
	}
	d.Words[eng] = vie
	return nil
}

func (d *Dictionary) UpdateWord(keyword Word, target Word) error {
	if status := d.HasWord(keyword); !status {
		return errors.New("word not found")
	}
	d.Words[keyword] = target
	return nil
}

func (d *Dictionary) DeleteWord(keyword Word) error {
	if status := d.HasWord(keyword); !status {
		return errors.New("word not found")
	}
	delete(d.Words, keyword)
	return nil
}

func (d *Dictionary) HasWord(keyword Word) bool {
	_, existing := d.Words[keyword]
	return existing
}
