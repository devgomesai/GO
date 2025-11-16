package main

import (
	"fmt"
	"strings"
)

func main() {
	var my_string = []rune("resume")
	var indexed = my_string[0]
	fmt.Printf("%v, %T\n\n", indexed, indexed)

	for i, v := range my_string {
		fmt.Println(i, v)
	}

	fmt.Printf("\nThe length of 'myString' is %v", len(my_string))

	strSlice := []string{"j", "o", "n", "a", "t", "h", "a", "n"}
	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i] + "_")
	}
	catStr := strBuilder.String()
	fmt.Printf("\n%v", catStr)
}
