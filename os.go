package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("args\t: %v\n", args)
	for _, arg := range args {
		fmt.Printf("arg: %s\n", arg)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("hostname:", hostname)
	}
}
