package service

import (
	"Day_3_-_Worker_Management/model"
	"Day_3_-_Worker_Management/utils"
	"container/list"
	"errors"
)

var listOfWorkers = list.New()

func CreateWorker() error {
	name, _ := utils.GetStringByRegex("Enter name: ", "", "")
	workLocation, _ := utils.GetStringByRegex("Enter location: ", "", "")
	age, _ := utils.GetInt("Enter age: ", "Age must be between 18 and 50", 18, 50)
	salary, _ := utils.GetDouble("Enter salary: ", "Salary must be at least 10 million VND", 1000000, 100000000)

	worker := model.NewWorker(name, age, salary, workLocation)

	if HasWorker(worker) {
		return errors.New("this worker is existing")
	}

	listOfWorkers.PushBack(worker)
	return nil
}

// signal true for increasing, false for decreasing
func UpdateSalary(signal bool) error {
	ID, _ := utils.GetStringByRegex("Enter ID: ", "", "")
	amount, _ := utils.GetDouble("Enter amount: ", "Salary must be larger than 0", 0, 100000000)
	worker := FindWorker(ID)
	if worker != nil {
		oldSalary := worker.GetSalary()
		if signal {
			worker.SetSalary(oldSalary + amount)
		} else {
			if oldSalary < amount {
				return errors.New("amount must be smaller than amount")
			}
			worker.SetSalary(oldSalary - amount)
		}
	}
	return nil
}

func HasWorker(w model.Worker) bool {
	for e := listOfWorkers.Front(); e != nil; e = e.Next() {
		worker := e.Value.(model.Worker)
		if w.GetName() == worker.GetName() && w.GetAge() == worker.GetAge() && w.GetWorkLocation() == worker.GetWorkLocation() {
			return true
		}
	}
	return false
}

func FindWorker(id string) *model.Worker {
	for e := listOfWorkers.Front(); e != nil; e = e.Next() {
		worker := e.Value.(*model.Worker)
		if worker.GetID().String() == id {
			return worker
		}
	}
	return nil
}
