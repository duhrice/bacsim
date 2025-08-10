package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Variables
	// Integers
	var intNum int = 22
	intNum = intNum + 1
	fmt.Println(intNum)

	// Floats
	var floatNum float64 = 22.22
	fmt.Println(floatNum)
	// When using arithmetic operators, do not mix data types
	// If you need to, cast it into the desired type
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	// Integer division will give a round down integer
	// Use modulo to get remainder
	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	// Strings (different from runes, uses double quotes)
	// Use double quotes to have a single-lined string (can not be broken across lines)
	// Use a single back quote to directly format string (can be broken across lines)
	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	// Built-in len() returns the number of bytes
	fmt.Println(len("test"))
	// RuneCountInString returns the number of characters in the string
	fmt.Println(utf8.RuneCountInString("test"))

	// Runes (different from strings, uses single quotes)
	var myRune rune = 'a'
	fmt.Println(myRune)

	// Booleans
	var myBoolean bool = false
	fmt.Println(myBoolean)

	// Declaring a variable without initializing it will set it to a default according to its data type
	var intNum3 int
	fmt.Println(intNum3)

	// Type does not have to be delcared if it is immediately initialized since it will be infered
	// var keyword can also be dropped when using the shorthand :=
	// Generally, its good practice to use var keyword and declare what data type the variable is
	var myVar1 = "text"
	myVar2 := "text"
	fmt.Println(myVar1, myVar2)

	// Multiple variables can be initialized at once
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// Constants
	// Must be declared and initialized with a value
	const myConst string = "const value"
	fmt.Println(myConst)
	// go run .\go_tutorials\tutorial_2\main.go
}