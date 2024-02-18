package main

import (
	"fmt"
	"go-std-libraries/models"
	"sort"
)

func main() {
	users := []models.User{
		{
			Name: "Thoriq",
			Age:  21,
		},
		{
			Name: "Fajra Risqulla",
			Age:  20,
		},
		{
			Name: "Cisnux",
			Age:  21,
		},
	}
	sort.Sort(models.UserSlice(users))
	fmt.Println(users)
}
