package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int32 = 10
	var y int64 = int64(x)

	fmt.Println(y)

	var z float64 = float64(y)

	fmt.Println(z)

	var b float64 = 3.3
	var a int32 = int32(b)
	fmt.Println(a)

	var grade_str string = "100"
	grade_int, _ := strconv.Atoi(grade_str) // string to integer

	fmt.Println(grade_int)

	stringValue := strconv.Itoa(-100) // integer to string
	fmt.Println(stringValue)

}
