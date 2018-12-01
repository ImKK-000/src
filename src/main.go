package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type Model struct {
	Name   string `validate:"required"`
	Number string `validate:"required"`
}

func main() {
	validate := validator.New()
	model := Model{
		Name:   "",
		Number: "",
	}
	err := validate.Struct(model)
	if err != nil {
		fmt.Println(err)
	}
}
