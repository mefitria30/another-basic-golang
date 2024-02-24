package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Test 0"
	names[1] = "Test 1"
	names[2] = "Test 2"

	fmt.Println(names)
	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Println("index", index, "=", value)
	}

	for _, value := range names {
		fmt.Println(value)
	}
}
