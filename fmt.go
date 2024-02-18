package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	firstName := "Fajra"
	lastName := "Risqulla"

	fmt.Printf("Hello %s %s!\n", firstName, lastName)
	fmt.Printf("memory address of firstName\t:%p\n", &firstName)
	fmt.Printf("memory address of lastName\t:%p\n", &lastName)
}
