package main

import "fmt"

// creating own data model type
type gasEngine struct {
	mpg     uint8
	gallons uint8
	// OwnerInfo owner
}

var _ engine = (*gasEngine)(nil)
var _ engine = (*electricEngine)(nil)

type electricEngine struct {
	mpkwh uint64
	kwh   uint64
}

// type owner struct {
// 	name string
// }

// Engine 1
func (e gasEngine) milesLeft() uint16 {
	return uint16(e.gallons) * uint16(e.mpg)
}

// Engine 2
func (e electricEngine) milesLeft() uint16 {
	return uint16(e.kwh) * uint16(e.mpkwh)
}

func (e gasEngine) knw() uint {
	return uint(e.gallons) * uint(e.mpg)
}

// Engine 2
func (e electricEngine) knw() uint {
	return uint(e.kwh) * uint(e.mpkwh)
}

// define an interface (the engine's must have the milesLeft() method)
type engine interface {
	milesLeft() uint16
	knw() uint
}

func canMakeIt(e engine, miles uint16) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it! ")
	} else {
		fmt.Println("Need to fuel up first! ")
	}
}

func main() {
	my_engine := gasEngine{mpg: 25, gallons: 50}
	// println(my_engine.gallons, my_engine.mpg)
	canMakeIt(my_engine, 70)

	var eletric_engine electricEngine = electricEngine{mpkwh: 7, kwh: 2}
	canMakeIt(eletric_engine, 100)
	fmt.Printf("Total miles left in tank: %v\n", my_engine.milesLeft())
	// my_arr := []gasEngine{
	// 	{mpg: 25, gallons: 50, OwnerInfo: owner{"Jonathan"}},
	// 	{mpg: 40, gallons: 35, OwnerInfo: owner{"Juanita"}},
	// }
	print()
	// for _, gase := range my_arr {
	// 	println(gase.mpg, gase.gallons, gase.OwnerInfo.name)
	// }

}
