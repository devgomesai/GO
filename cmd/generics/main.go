package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(len(intSlice))
	fmt.Println(sumSlice(intSlice))

	// var intSlice1 = []float64{1.1, 2, 3}
	// fmt.Println(sumSlice(intSlice1))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
