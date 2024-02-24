package main

import "fmt"

func main() {
	for i := 1; i <= 30; i++ {
		// if i%2 == 1 {
		// 	fmt.Println(i, "Odd")
		// } else {
		// 	fmt.Println(i, "Even")
		// }

		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FooBar")
		} else if i%3 == 0 {
			fmt.Println(i, "Foo")
		} else if i%5 == 0 {
			fmt.Println(i, "Bar")
		} else {
			fmt.Println(i)
		}
	}
}
