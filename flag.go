package main

import (
	"flag"
	"fmt"
)

func main() {
	id := flag.Int("id", 1, "Enter your id")
	host := flag.String("host", "localhost", "Enter your database host")
	username := flag.String("username", "root", "Enter your database username")
	password := flag.String("password", "root", "Enter your database password")

	flag.Parse()

	fmt.Printf("id\t\t: %d\n", *id)
	fmt.Printf("host\t\t: %s\n", *host)
	fmt.Printf("username\t: %s\n", *username)
	fmt.Printf("password\t: %s\n", *password)
}
