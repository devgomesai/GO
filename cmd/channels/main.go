package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	var c = make(chan int)
// 	go process(c)
// 	for i := range c {
// 		fmt.Println(i)
// 		time.Sleep(time.Second * 1)
// 	}

// }

//	func process(c chan int) {
//		defer close(c) // close before the channel exists
//		for i := 0; i < 5; i++ {
//			c <- i
//		}
//		fmt.Println("Exiting process")
//	}
var MAX_CHICKEN_PRICE float32 = 5

func main() {
	var chickenchannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenchannel)
	}
	sendMessage(chickenchannel)
}

func checkChickenPrices(website string, chickenchannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		fmt.Println(chickenPrice, " :> ", website)
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenchannel <- website
			break
		}
	}
}

func sendMessage(channel chan string) {
	fmt.Printf("Found a deal on chicken at %s", <-channel)
}
