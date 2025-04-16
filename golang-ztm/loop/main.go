package main

import (
	"math/rand"
)

// func main() {

// 	for i := 1; i < 50; i++ {
// 		divisibleBy3 := i%3 == 0
// 		divisibleBy5 := i%5 == 0
// 		if divisibleBy3 && divisibleBy5 {
// 			fmt.Println("FizzBuzz")
// 		}
// 		if divisibleBy3 && !divisibleBy5 {
// 			fmt.Println("Fizz")
// 		}
// 		if !divisibleBy3 && divisibleBy5 {
// 			fmt.Println("Buzz")
// 		}
// 		if !divisibleBy3 && !divisibleBy5 {
// 			fmt.Println(i)
// 		}
// 	}
// }

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	// dice, sides := 2, 12
	// rolls := 1

}
