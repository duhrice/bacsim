package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func process(c chan int) {
	// Close the channel to indicate that no more values are being passed through the channel
	defer close(c)
	for i := 0; i < 5; i++ {
		// Use <- to pass information into and out of the channel
		c <- i
	}
	fmt.Println("Exiting process")
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	// Select essentially acts as an if for channels
	select {
		case website := <- chickenChannel:
			fmt.Printf("Found a deal on chicken at %v\n", website)
		case website := <- tofuChannel:
			fmt.Printf("Email sent: Found a deal on chicken at %v\n", website)
	}
}

func main() {
	// Channels (pipelines that allow information to be sent efficiently between go rountines)
	// Only holds 1 value
	// var  c = make(chan int)
	// Buffer channel to hold more values
	var  c = make(chan int, 5)
	go process(c)
	for i := range c{
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}

	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string {"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
	// go run .\go_tutorials\tutorial_9\main.go
}