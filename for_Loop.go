package main

import (
	"fmt"
)

func main() {
	// //simple iteration over a range
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// //iterate over a collection
	// numbers := []int {1,2,3,4,5,6}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i % 2 == 0 {
	// 		continue //skip even numbers
	// 	}
	// 	fmt.Println("Odd Number:", i)
	// 	if i == 5 {
	// 		break //break out of the loop when i is 5
	// 	}
	// }

	// rows := 5
	// //outer loop
	// for i := 1; i <= rows; i++ {
	// 	//inner loop
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	//inner loop for stars
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("+")
	// 	}
	// 	fmt.Println()
	// }

	for i := range 10 {
		fmt.Println(i)
	}
	fmt.Println("We have lift off!")
}