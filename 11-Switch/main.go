package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 1; i <= 10; i++ {
		switch i {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Undefined")
		}
	}

	OperatingSystem := runtime.GOOS

	switch OperatingSystem {
	case "darwin":
		fmt.Println("Max")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Undefined")
	}

}
