package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "name,age,email\n" +
		"fajra,20,fajra@gmail.com\n" +
		"cisnux,18,cisnux@gmail.com\n" +
		"risqulla,19,risqulla@gmail.com"

	var stringReader io.Reader = strings.NewReader(csvString)
	var reader *csv.Reader = csv.NewReader(stringReader)
	for record, err := reader.Read(); err != io.EOF; record, err = reader.Read() {
		fmt.Println(record)
	}
	//for {
	//	record, err := reader.Read()
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(record)
	//}
}
