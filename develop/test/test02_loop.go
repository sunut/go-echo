package main

import (
	"fmt"
)

func main() {
	sum := 0
	fmt.Println("test loop just use syntax similar C lang but not `(` or `)` ")
	/*this is for loop*/
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("total sum in loop 10 times(0-9) is", sum)

	/*this is for while loop*/
	fmt.Println("test while loop : add condition in for loop")
	sum = 0
	for sum < 1000 {
		sum += 100
	}
	fmt.Println("total sum in loop", sum)

	/*while loop (true condition)
	for {
		if (condition){
			break
		}
	}
	*/

}
