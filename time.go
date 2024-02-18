package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2002, time.June, 12, 03, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	value := "2022-02-25 02:20:20"
	formatter := "2006-01-02 15:04:05"
	//formatter := time.RFC3339
	parsed, err := time.Parse(formatter, value)
	if err != nil {
		println("Error", err.Error())
	} else {
		fmt.Println(parsed)
	}
}
