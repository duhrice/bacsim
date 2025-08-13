package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

// Method (notice that the parameters are before the function name)
// Output type is the method signature
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

// Interface (similar to making a struct but can be used more generally)
// Any methods that function similarly and return the required method signature is ok
type engine interface {
	milesLeft() uint8
}

// Instead of "e gasEngine" or "e electicEngine", "e" can now be any type of engine
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
} 

func main() {
	// Declaring a variable without initialization will use default values for the types in the struct
	var myEngine gasEngine = gasEngine{mpg:25, gallons:15} // or {25, 15}
	var myEngine2 electricEngine
	// This type of anonymous struct is not reusable
	// var myEngine2 = struct {
	// 	mpg.uint8
	// 	gallons.uint8
	// } {25, 15}
	myEngine.mpg = 20
	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())
	canMakeIt(myEngine, 50)
	canMakeIt(myEngine2, 50)
	// go run .\go_tutorials\tutorial_6\main.go
}