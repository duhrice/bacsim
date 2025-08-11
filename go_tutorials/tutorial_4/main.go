package main

import "fmt"

func main() {
	// Arrays
	// Only store one data type
	// Uses the default value of that data type when declared but not initialized
	// Has a fixed size when initialized
	// var intArr [3]int32 = [3]int32 {1, 2, 3}

	// Size of the array can be inferred by using elipses
	intArr := [...]int32 {1, 2, 3}
	fmt.Println(intArr)
	intArr[1] = 123
	fmt.Println(intArr[0])
	// Access element 1 and 2 (start from index i, go up to but not including index j)
	fmt.Println(intArr[1:3])

	// Get the memory location using &
	// Values are stored in continguous memory
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// Slices (arrays under the hood but more powerful because it has some added functionality)
	var intSlice []int32 = []int32 {4, 5, 6}
	fmt.Printf("The length is %v with capcity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capcity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32 {8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// Another way o fmaking slices
	// Optional argument for length and capacity
	// Setting a capacity ahead of time is faster
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	// Maps
	// Maps always return a value
	// Key data type in racket, value data type outside bracket
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8 {"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	// Using a key not in the map returns the default value of the value data type
	// fmt.Println(myMap2["Jason"])
	// Maps can return an optional boolean as a second value
	// True if in map, False otherwise
	var age, ok = myMap2["Jason"]
	delete(myMap2, "Adam")
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid name")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	// No while loops, use for loops only
	var i int = 0
	// for i := 0; i < 10; i++
	for i < 10 {
		/*
		if i >= 10 {
			break
		}
		*/
		fmt.Println(i)
		i = i + 1
	}
	// go run .\go_tutorials\tutorial_4\main.go
}