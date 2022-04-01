package main

import (
	"fmt"
)

func main() {
	fmt.Println("test pointer")
	valueA := "hello"
	fmt.Println(valueA)
	var valueB *string = &valueA
	fmt.Println(valueB)
	*valueB = "new hello"
	fmt.Println(*valueB)
	fmt.Println(valueA)
}
