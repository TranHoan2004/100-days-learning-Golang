package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func RemoveUnnecessaryBlank(input string) string {
	re := regexp.MustCompile("\\s+")
	trimmed := strings.TrimSpace(input)
	return re.ReplaceAllString(trimmed, " ")
}

func GetInt(message string, error string, min int, max int) (int, error) {
	input, err := GetStringByRegex(message, "[0-9]*\\\\.?[0-9]+", error)
	if err != nil {
		return 0, err
	}

	number, err1 := strconv.Atoi(input)
	if err1 != nil {
		return 0, err1
	}

	if number < min || number > max {
		return 0, errors.New("out of range")
	}
	return number, nil
}

func GetDouble(message string, error string, min float32, max float32) (float32, error) {
	input, err := GetStringByRegex(message, "[0-9]+", error)
	if err != nil {
		return 0, err
	}

	numberRaw, err1 := strconv.ParseFloat(input, 32)
	if err1 != nil {
		return 0, err1
	}

	number := float32(numberRaw)

	if number < min || number > max {
		return 0, errors.New("out of range")
	}
	return number, nil
}

func GetStringByRegex(message string, regex string, error string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(message)
	var text string
	if scanner.Scan() {
		text = RemoveUnnecessaryBlank(scanner.Text())
		if text == "" {
			return "", errors.New("not null")
		}
		match, err := regexp.MatchString(regex, text)
		if err != nil {
			return "", errors.New(fmt.Sprintf("Invalid regexp: %v", err))
		}
		if !match {
			return "", errors.New(error)
		}
	}
	return text, nil
}
