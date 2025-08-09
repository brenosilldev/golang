package main

import "fmt"

// Closures uma função dentro de outra função.

func main() {

	total := func() int {
		return 10
	}()

	fmt.Println(total)

}
