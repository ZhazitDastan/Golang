package main

import "fmt"

func main() {

	//The result will not change. To make the output look like slice, I added the character '[', ']'.
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := []int{90, 80, 50, 223, 10, 15, 43, 49, 51, 36}

	PrintSlice(slice1)
	PrintSlice(slice2)
}

func PrintSlice(slice []int) {
	fmt.Print("[")

	for i, value := range slice {
		fmt.Print(value)

		if i < len(slice)-1 {
			fmt.Print(", ")
		}
	}

	fmt.Println("]")
}
