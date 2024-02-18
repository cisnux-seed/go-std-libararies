package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Fajra Risqulla", "Fajra"))
	fmt.Println(strings.Split("Fajra Risqulla", " "))
	fmt.Println(strings.Join([]string{"Fajra", "Risqulla"}, ","))
	fmt.Println(strings.ToLower("Fajra Risqulla"))
	fmt.Println(strings.Trim("		Fajra Risqulla		", "		"))
	fmt.Println(strings.ReplaceAll("Fajra Fajra Fajra Fajra", "Fajra", "Cisnux"))
}
