package main

import "fmt"

// type gasEngine struct {
// 	gallons float32
// 	mpg float32
// }

// type electricEngine struct {
// 	kwh float32
// 	mpkeh float32
// }

// type car [T gasEngine | electricEngine] struct {
// 	carMake string
// 	carModel string
// 	engine T
// }

// Generic (allow the acceptance of multiple types and returns values of the specificed type)
func sumSlice[T int | float32 | float64] (slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any] (slice []T) bool {
	return len(slice) == 0
}

func main() {
	var intSlice = []int {1, 2, 3}
	// You can specify the type (also includes structs) but go can infer it (case by case basis)
	// fmt.Println(sumSlice[int](intSlice))
	fmt.Println(sumSlice(intSlice))
	fmt.Println(isEmpty(intSlice))

	var float32Slice = []float32 {1, 2, 3}
	fmt.Println(sumSlice(float32Slice))
	fmt.Println(isEmpty(float32Slice))

	// var gasCar = car[gasEngine] {
	// 	carMake: "Honda",
	// 	carModel: "Civic",
	// 	engine: gasEngine{
	// 		gallons: 12.4,
	// 		mpg: 40,
	// 	},
	// }
	// var electricCar = car[electricEngine] {
	// 	carMake: "Tesla",
	// 	carModel: "Model 3",
	// 	engine: electricEngine{
	// 		kwh: 57.4,
	// 		mpkeh: 4.17,
	// 	},
	// }

	// go run .\go_tutorials\tutorial_10\main.go
}