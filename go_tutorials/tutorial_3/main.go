package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is %v", result)
	// } else {
	// 	fmt.Printf("The result of the inteder division is %v with remainder %v", result, remainder)
	// }

	// Switch cases
	switch {
		case err != nil:
			fmt.Printf("%s", err.Error())
		case remainder == 0:
			fmt.Printf("The result of the integer division is %v\n", result)
		default:
			fmt.Printf("The result of the inteder division is %v with remainder %v\n", result, remainder)
	}

	// Conditional switch cases
	switch remainder {
		case 0:
			fmt.Printf("The division was exact")
		case 1, 2:
			fmt.Printf("The division was close")
		default:
			fmt.Printf("The division was not close")
	}
	// go run .\go_tutorials\tutorial_3\main.go
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	// Default value is nil
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}