package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const fileName = "test.text"

func main() {
	err := createNewFile(fileName, "Hello World\n")
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := 0; i < 10; i++ {
		err := addNewData(fileName, fmt.Sprintf("Hello World %d\n", i))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	message, err := readFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Print(message)
	}
}

func createNewFile(name string, message string) (err error) {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			return
		}
	}()
	_, err = file.WriteString(message)
	return
}

func readFile(name string) (message string, err error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			return
		}
	}()
	reader := bufio.NewReader(file)
	for line, _, err := reader.ReadLine(); err != io.EOF; line, _, err = reader.ReadLine() {
		message += fmt.Sprintf("%s\n", string(line))
	}
	return
}

func addNewData(name string, message string) (err error) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			return
		}
	}()
	_, err = file.WriteString(message)
	return
}
