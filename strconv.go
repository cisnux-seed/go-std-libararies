package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(boolean)
	}
	integer64, err := strconv.ParseInt("2938382829", 10, 64)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(integer64)
	}
	fmt.Println(strconv.FormatFloat(2.89, 'g', -1, 32))
}
