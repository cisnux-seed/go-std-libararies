package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	names := []string{"risqulla", "cisnux", "jack", "fajra risqulla"}
	numbers := []int{100, 75, 80, 20}
	fmt.Println(slices.Min(numbers))
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(numbers))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Contains(names, "fajra"))
	fmt.Println(slices.Contains(numbers, 75))
	fmt.Println(slices.ContainsFunc(names, func(s string) bool {
		return strings.Contains(s, "fajra")
	}))
	fmt.Println(slices.Index(numbers, 80))
	fmt.Println(slices.Index(names, "fajra risqulla"))
}
