package main

import (
	"github.com/alimy/dugo/learn"
	"fmt"
)

func main() {
	fmt.Println("============= hello world ================")
	learn.HelloWorld()

	fmt.Println("============= slice ================")
	learn.TrySlice()

	fmt.Println("============= slice remove ================")
	learn.TrySliceRemove()
	learn.TrySliceRemove2()

	fmt.Println("============= slice nonEmpty ================")
	learn.TrySliceNonEmpty()
	learn.TrySliceNonEmpty2()

	fmt.Println("============= map ================")
	learn.TryMap()

	fmt.Println("============= map delete================")
	learn.TryMapDelete()
}


