package main

import "fmt"

func main() {
	// var p *int32 = new(int32)
	// var f *int32 = new(int32)
	// var i int32
	// fmt.Printf("The value p points to is : %v", *p)
	// fmt.Printf("\nThe address of p is : %p", &p)
	// fmt.Printf("\nThe value of i is : %v", i)

	// *p = 10
	// fmt.Print("\n")
	// f = p

	// fmt.Printf("\nf points to address: %p", f)
	// fmt.Printf("\nThe value p points to is : %v", *p)
	// fmt.Printf("\nThe value f points to is : %v\n", *f)

	var thing1 = [...]float64{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Initial Array: %v", thing1)
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	var result [7]float64 = square(&thing1)
	fmt.Printf("\nResult: %v", result)
}

func square(things *[7]float64) [7]float64 {
	fmt.Printf("\nThe  memory location of things is: %p", things)

	for i := range things {
		things[i] = things[i] * things[i]
	}

	return *things
}
