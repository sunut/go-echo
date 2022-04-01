package main

import (
	"fmt"
	"math/rand"
)

func calculate(number int) bool {
	if number < 100 {
		return true
	} else {
		return false
	}
}

func calculateRandom(number int) bool {
	/*you can create variable scope use in condition*/
	if randomNumber := rand.Intn(100); number < randomNumber {
		return true
	} else {
		return false
	}
}

func calculateCompare2Value(number1, number2 int) (bool, string) {
	/*1.you can  asign number1 & number2  to int
	  2.a function can return multiple value
	*/
	if number1 < number2 {
		return true, "number2 more than number1"
	} else {
		return false, "number1 more than number2"
	}
}

func calculateTestFunction() (x int, y int, result string) {
	x = 1
	y = 1
	if x < y {
		return x, y, "number2 more than number1"
	} else {
		return
	}
}

func main() {
	fmt.Println("test condition")
	fmt.Println(calculate(99))
	fmt.Println(calculate(10000))
	fmt.Println(calculateRandom(10000))
	fmt.Println(calculateCompare2Value(10, 1000))
	fmt.Println(calculateTestFunction())
}
