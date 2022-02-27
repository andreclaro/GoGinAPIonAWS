package main

import (
	"errors"
	"log"
	"strconv"
)

func convertToInteger(s string) (int, error) {
	number, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)
		return 0, errors.New("can't work with 42")
	}
	return number, nil
}

func calculate(valueA int, valueB int, operator string) int {
	switch operator {
	case "+":
		return valueA + valueB
	case "=":
		return valueA - valueB
	case "*":
		return valueA * valueB
	case "/":
		return valueA / valueB
	}
	return valueA + valueB
}
