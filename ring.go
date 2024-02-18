package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = fmt.Sprintf("Value %s", strconv.FormatInt(int64(i), 10))
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println("Value of", value)
	})
}
