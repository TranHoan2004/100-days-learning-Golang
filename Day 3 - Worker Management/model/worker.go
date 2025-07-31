package model

import (
	"fmt"
	"github.com/google/uuid"
)

type Worker struct {
	ID           uuid.UUID
	Name         string
	Age          int
	Salary       float32
	WorkLocation string
}

func NewWorker(name string, age int, salary float32, workLocation string) Worker {
	return Worker{
		ID:           uuid.New(),
		Name:         name,
		Age:          age,
		Salary:       salary,
		WorkLocation: workLocation,
	}
}

func (w *Worker) SetID(ID uuid.UUID) {
	w.ID = ID
}

func (w *Worker) GetID() uuid.UUID {
	return w.ID
}

func (w *Worker) SetName(Name string) {
	w.Name = Name
}

func (w *Worker) GetName() string {
	return w.Name
}

func (w *Worker) SetAge(Age int) {
	w.Age = Age
}

func (w *Worker) GetAge() int {
	return w.Age
}

func (w *Worker) SetSalary(Salary float32) {
	w.Salary = Salary
}

func (w *Worker) GetSalary() float32 {
	return w.Salary
}

func (w *Worker) SetWorkLocation(Location string) {
	w.WorkLocation = Location
}

func (w *Worker) GetWorkLocation() string {
	return w.WorkLocation
}

func (w *Worker) ToString() string {
	return fmt.Sprintf("Worker:\n"+
		"ID: %v\n"+
		"Name: %v\n"+
		"Age: %v\n"+
		"Salary: %v\n"+
		"Work Location: %v",
		w.ID, w.Name, w.Age, w.Salary, w.WorkLocation)
}
