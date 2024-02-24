package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i, "Hello World")
		i++
	}

	fmt.Println("End")

	// or the normal look

	for x := 1; x <= 10; x++ {
		fmt.Println(x, "Hello World")

	}

	fmt.Println("End")

	// decrement

	for y := 10; y >= 1; y-- {
		fmt.Println(y, "Hello World")

	}

	fmt.Println("End")
}
