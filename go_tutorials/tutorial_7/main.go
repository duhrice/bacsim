package main

import "fmt"

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of the thing2 array is: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)
	p = &i
	*p = 1
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	var thing1 = [5]float64 {1, 2, 3, 4, 5}
	fmt.Printf("The memory location of the thing2 array is: %p\n", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("The result is: %v\n", result)
	fmt.Printf("The value of thing1 is: %v", thing1)
	// go run .\go_tutorials\tutorial_7\main.go
}