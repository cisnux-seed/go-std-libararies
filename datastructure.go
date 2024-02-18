package main

import (
	"fmt"
	"go-std-libraries/datastructures"
)

func main() {
	list := datastructures.New()
	list.PushBack("Hello")
	list.PushBack("World")
	list.PushBack(1)
	list.PushBack(5)
	for e := list.Front(); e != nil; e = list.Next() {
		fmt.Println(e.Value)
	}
	list.RemoveFirst()
	list.RemoveLast()
	list.PushBack(10)
	for e := list.Front(); e != nil; e = list.Next() {
		fmt.Println(e.Value)
	}
}
