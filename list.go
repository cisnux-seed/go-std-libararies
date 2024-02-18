package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Fajra")
	data.PushBack("Risqulla")
	data.PushBack("Cisnux")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
		fmt.Println(e)
	}
}
