package main

import "fmt"

func main() {
	a := 1
	b := 1

	result_a := a == b // a equal b
	result_b := a != b // a not equal b
	result_c := a < b  // a lower than b
	result_d := a > b  // a greater than b
	result_e := a <= b // a lower than equal b
	result_f := a >= b // a greater than equal b

	fmt.Println(result_a)
	fmt.Println(result_b)
	fmt.Println(result_c)
	fmt.Println(result_d)
	fmt.Println(result_e)
	fmt.Println(result_f)
}
