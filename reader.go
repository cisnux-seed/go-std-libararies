package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("this is log string\nwith new line\n")
	reader := bufio.NewReader(input)

	for line, _, err := reader.ReadLine(); err != io.EOF; line, _, err = reader.ReadLine() {
		fmt.Println(string(line))
	}
}
