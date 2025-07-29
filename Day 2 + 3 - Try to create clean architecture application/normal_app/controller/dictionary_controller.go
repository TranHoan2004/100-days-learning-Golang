package controller

import (
	"bufio"
	"fmt"
	"normal_app/service"
	"os"
)

type DictionaryController struct {
	service *service.Service
}

func NewDictionaryController() *DictionaryController {
	return &DictionaryController{
		service: service.NewService(),
	}
}

func (c *DictionaryController) CreateNewWord() string {
	engWord := GetWord("Enter english word: ")
	vieWord := GetWord("Enter vietnamese word: ")
	err := c.service.AddWord(engWord, vieWord)
	if err != nil {
		return err.Error()
	}
	return "Add words successfully"
}

func (c *DictionaryController) UpdateWord() string {
	target := GetWord("Enter the English word that you want to update: ")
	content := GetWord("Enter new Vietnamese word: ")
	return c.service.UpdateWord(target, content)
}

func (c *DictionaryController) DeleteWord() string {
	target := GetWord("Enter the English word that you want to delete: ")
	return c.service.DeleteWord(target)
}

func (c *DictionaryController) DisplayDictionaryContent() {
	fmt.Println("Content of dictionary")
	for eng, vie := range c.service.GetWords() {
		fmt.Printf("English: %v, Vietnamese: %v\n", eng, vie)
	}
}

func GetWord(title string) string {
	scanner := bufio.NewScanner(os.Stdin)
	var eng string
	fmt.Print(title)
	if scanner.Scan() {
		eng = scanner.Text()
	}
	return eng
}
