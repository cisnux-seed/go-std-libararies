package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`f([a-z])jra`)

	fmt.Println(regex.MatchString("fajra"))
	fmt.Println(regex.MatchString("fAjra"))
	fmt.Println(regex.MatchString("fsjra"))
	// n=max
	fmt.Println(regex.FindAllString("fajra fAjra fsjra fajra f1jra fajra", -1))
	fmt.Println(regex.FindAllString("fajra fAjra fsjra fajra f1jra fajra", 2))
	fmt.Println(regex.FindAllString("fajra fAjra fsjra fajra f1jra fajra", 10))
}
