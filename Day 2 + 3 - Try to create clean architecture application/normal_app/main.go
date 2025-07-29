package main

import (
	"bufio"
	"fmt"
	"normal_app/controller"
	"os"
	"strconv"
)

func main() {
	c := controller.NewDictionaryController()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("=============== Menu ===============")
		fmt.Print("Enter option: ")
		var option int
		if scanner.Scan() {
			text := scanner.Text()
			order, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			option = order
		}
		switch option {
		case 1:
			{
				fmt.Println(c.CreateNewWord())
				break
			}
		case 2:
			{
				fmt.Println(c.UpdateWord())
				break
			}
		case 3:
			{
				fmt.Println(c.DeleteWord())
				break
			}
		case 4:
			{
				c.DisplayDictionaryContent()
				break
			}
		default:
			return
		}
	}
}
