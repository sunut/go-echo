package main

import (
	"fmt"
)

func main() {
	var supanut, nattakit bool
	fmt.Println("01 declar variable")
	fmt.Println(supanut, nattakit)
	var scb int
	var techx float64
	var aiscb bool
	var scbs string
	fmt.Println("02 default variable ")
	fmt.Println("int:->", scb, "float:->", techx, "bool:->", aiscb, "string:->", scbs)

	shouldBeString := "hello world"
	fmt.Println("03 assign variable")
	fmt.Println(shouldBeString)

}
