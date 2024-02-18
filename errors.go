package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func main() {
	err := GetById("csx")
	if err != nil {
		switch {
		case errors.Is(err, ValidationError):
			fmt.Println(err.Error())
		case errors.Is(err, NotFoundError):
			fmt.Println(err.Error())
		default:
			fmt.Println("unknown error")
		}
	}
}

func GetById(id string) (e error) {
	if id == "" {
		e = ValidationError
		return
	}

	if id != "cisnux" {
		e = NotFoundError
		return
	}
	return
}
