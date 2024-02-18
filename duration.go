package main

import (
	"fmt"
	"time"
)

func main() {
	duration1 := time.Second * 100
	duration2 := time.Microsecond * 10
	duration := duration1 - duration2

	fmt.Printf("duration: %d\n", duration)
}
