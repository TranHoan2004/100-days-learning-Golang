package controller

import (
	"Day_3_-_Worker_Management/model"
	"Day_3_-_Worker_Management/service"
	"fmt"
)

func AddNewWorker() {
	err := service.CreateWorker()
	if err != nil {
		fmt.Println(err)
	}
}

func ChangeSalaryAmount(signal bool) {
	err := service.UpdateSalary(signal)
	if err != nil {
		fmt.Println(err)
	}
}

func DisplayAllWorkers() {
	workersList := service.GetAllWorkers()

	for e := workersList.Front(); e != nil; e = e.Next() {
		worker := e.Value.(*model.Worker)
		fmt.Println(worker.ToString())
	}
}
