package main

import (
	"fmt"
	"go-std-libraries/models"
	"go-std-libraries/utils"
	"reflect"
)

func main() {
	sample := models.Sample{
		Name: "Fajra",
		Age:  20,
	}
	user := models.User{
		Name: "Cisnux",
		Age:  100,
	}
	user2 := models.User{
		Name: "fajra",
		Age:  101,
	}

	sampleType := reflect.TypeOf(sample)
	userType := reflect.TypeOf(user)

	fmt.Println("type of sample", sampleType.Name())
	for i := 0; i < sampleType.NumField(); i++ {
		field := sampleType.Field(i)
		fmt.Printf("%d name of field\t:%s\n", i+1, field.Name)
		fmt.Printf("%d type of field\t:%s\n", i+1, field.Type.Name())
		fmt.Printf("%d tag of field\t:%s\n", i+1, field.Tag.Get("required"))
		fmt.Printf("%d tag of field\t:%s\n", i+1, field.Tag.Get("max"))
	}
	println()
	fmt.Println("type of user", userType.Name())
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		fmt.Printf("%d name of field\t:%s\n", i+1, field.Name)
		fmt.Printf("%d type of field\t:%s\n", i+1, field.Type)
		fmt.Printf("%d tag of field\t:%s\n", i+1, field.Tag.Get("required"))
		fmt.Printf("%d tag of field\t:%s\n", i+1, field.Tag.Get("max"))
	}
	isValid, err := utils.IsValid(user)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("valid: ", isValid)
	}
	isValid2, err2 := utils.IsValid(user2)
	if err2 != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("valid: ", isValid2)
	}
}
