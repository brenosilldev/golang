package main

import "fmt"

func main() {

	for i, index := range []string{"a", "b", "c"} {
		fmt.Println("index:",index)
		fmt.Println("i:", i)
		fmt.Println()
		for _, cbvvd := range index {
			fmt.Println(cbvvd)
		}
	}

}
