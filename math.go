package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40))
	fmt.Println(math.Floor(1.60))
	fmt.Println(math.Round(1.60))
	fmt.Println(math.Round(1.49))
	fmt.Println(math.Max(1.50, 1.51))
	fmt.Println(math.Min(1.50, 1.51))
}
