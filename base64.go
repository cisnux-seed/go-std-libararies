package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encoded := base64.StdEncoding.EncodeToString([]byte("Fajra Risqulla Cisnux"))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decoded))
		for _, b := range decoded {
			fmt.Println(b)
		}
	}
	var b byte = 'A'
	fmt.Println(b)
	fmt.Println(string(b))
}
