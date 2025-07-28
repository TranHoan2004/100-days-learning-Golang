package model

import (
	"fmt"
	"github.com/google/uuid"
)

// go get github.com/google/uuid

type Word struct {
	ID   uuid.UUID
	Name string
}

func NewWord(name string) *Word {
	return &Word{
		ID:   uuid.New(),
		Name: name,
	}
}

func (w *Word) GetID() uuid.UUID {
	return w.ID
}

func (w *Word) GetName() string {
	return w.Name
}

func (w *Word) SetID() {
	w.ID = uuid.New()
}

func (w *Word) SetName(name string) {
	w.Name = name
}

func (w *Word) ToString() string {
	return fmt.Sprintf("Word: %v, %v", w.ID.String(), w.Name)
}
