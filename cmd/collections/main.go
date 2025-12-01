package main

import "fmt"

func main() {
	// define the array
	var intArr [10]int32
	// func
	for i := 0; i < 10; i++ {
		intArr[i] = int32(i * 1)
	}
	// see the array
	//fmt.Println(intArr) // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

	// get address of that element stored
	//fmt.Println(&intArr[7]) // get the address

	// int2Arr := [5]int32{1, 2, 3, 4, 5}
	// int2Arr := []int32{}
	// fmt.Println("Initial Array: ", int2Arr)

	// intSlice := append(int2Arr, 7)
	// fmt.Printf("The length is %v with capacity %v: ", len(intSlice), cap(intSlice))

	// var intSlice2 []int32 = []int32{8, 9}
	// intSlice2 = append(intSlice, intSlice2...)
	// fmt.Println(intSlice2)

	// var intSlice3 []int32 = make([]int32, 3,8)
	// fmt.Println(cap(intSlice3))

	myMap := make(map[string]int)
	myMap["Jonathan"] = 21
	myMap["Reynold"] = 20
	myMap["Ranjeeta"] = 30

	for name, age := range myMap {
		fmt.Println("Name: ", name, "Age: ", age)
	}

	// fmt.Println(myMap["Jonathan"])

	// age, ok := myMap["Jonathan"]
	// fmt.Println(age)
	// fmt.Println(ok)
	fmt.Printf("Before Deletion: %v\n", myMap)
	delete(myMap, "Jonathan")
	fmt.Printf("After Deletion : %v\n", myMap)

	// for name, age := range myMap {
	// 	fmt.Printf("Name:%v, \nAge:%v \n", name, age)
	// }

	// var name, sname string
	// fmt.Print("Enter Name: ")
	// fmt.Scan(&name, &sname)

	// fmt.Println(name, sname)
	// var n int = 10
	// for i := 0; i <= n; i++ {
	// 	fmt.Println(i)
	// }

	myArray := [...]int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range myArray {
		println("index: ", i, " :> ", " value: ", v)
	}

}
