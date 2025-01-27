package main

import (
	"fmt"
)

func main() {
	numbers := []int{10, -10, 5}
	var S int
	for _, number := range numbers {
		if number > 0 {
			S = S + number
		}
	}
	fmt.Println(S)
}

//// No site do codewars ficou assim:

// package kata

// func PositiveSum(numbers []int) int {
//   var Soma int
//   for _, number := range numbers {
//     if number > 0 {
//       Soma = Soma + number
//     }
//   }
//     return Soma
// }
