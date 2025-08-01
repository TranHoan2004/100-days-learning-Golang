package main

import (
	"Day_3_-_Worker_Management/controller"
	"Day_3_-_Worker_Management/utils"
	"fmt"
)

func main() {
	for {
		option, err := GetOption()
		if err != "" {
			fmt.Println(err)
		}
		switch option {
		case 1:
			// Add a worker
			controller.AddNewWorker()
			break
		case 2:
			// Increase salary
			controller.ChangeSalaryAmount(true)
			break
		case 3:
			// Decrease salary
			controller.ChangeSalaryAmount(false)
			break
		case 4:
			// Get all
			controller.DisplayAllWorkers()
			break
		case 5:
			return
		default:
		}
	}
}

func GetOption() (int, string) {
	fmt.Println("================= Menu =================")
	option, err := utils.GetInt("Enter your option: ", "Option must be from 1 to 5", 1, 5)
	if err != nil {
		return 0, err.Error()
	}
	return option, ""
}
