package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func AddNumbers(a int, b int) int {
	result := a + b
	return result
}

func main() {
	var intNum = 32767 // intNum := 1516
	fmt.Println("Hello World")
	fmt.Println(intNum)

	var float_num float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = float_num + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)

	my_str := "Hello 😂 World"
	fmt.Println(len(my_str))
	string_len := utf8.RuneCountInString(my_str)
	fmt.Println("The Length of the String ( "+my_str+" ): ", string_len)
	num1 := 10
	num2 := 4
	q, r, err := intDivision(num1, num2)
	if err != nil {
		fmt.Println(err.Error())
	}
	println(q, r)

	if err == nil {
		fmt.Println("Error Free Code")
	} else {
		fmt.Println("Error")
	}
	a := 10
	b := 20
	ans := AddNumbers(a, b)

	fmt.Println("Sum of (", a, ",", b, ") :", ans)

}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error = nil
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	quotient := numerator / denominator
	remainder := numerator % denominator
	return quotient, remainder, err
}
