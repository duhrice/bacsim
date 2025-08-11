package main

import (
	"fmt"
	"strings"
)

func main() {
	// Underlying value of string is an array of bytes
	// var myString string = "résumé"
	var myString = []rune("résumé")
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("The length of 'myString' is %v\n", len(myString))

	var myRune = 'a'
	fmt.Printf("myRune = %v", myRune)

	// Strings are immutable
	// Concatenation like down below is actually creating a new string every time
	var strSlice = []string {"t", "e", "s", "t"}
	// var catStr = ""
	// Using the strings library fixes this by first creating a array to store characters
	var strBuilder strings.Builder
	for i := range strSlice {
		// catStr += strSlice[i]
		strBuilder.WriteString(strSlice[i])
	}
	// Characters are then joined here to create a string
	var catStr = strBuilder.String()
	fmt.Println(catStr)
	// go run .\go_tutorials\tutorial_5\main.go
}