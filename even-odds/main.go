package main

import (
	"fmt"
)

func main() {
	numbers := []int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if number % 2 == 0 {
			fmt.Printf("\n%v is even", number)
		} else {
			fmt.Printf("\n%v is odd", number)
		}
	}
}