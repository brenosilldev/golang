package main

import (
	"fmt"
)

func twoSum() (int, int) {
	return 1, 2
}

func five() int {
	return 5
}

func AddThre(a, b, c int) int {
	return a + b + c
}

func greetPerson(name string) {
	fmt.Println("Hello", name)
}

func hiTrehe() string {
	return "hi"
}
func main() {
	a, b := twoSum()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(five())
	fmt.Println(AddThre(1, 2, 3))
	greetPerson("Breno")
	fmt.Println(hiTrehe())
}
