package main

import "fmt"

func main() {

	// result#1: 9 8 7 6 5 4 3 2 1
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// result#2: 10 20 30 40 50 60 70 80 90
	slice2 := []int{90, 80, 70, 60, 50, 40, 30, 20, 10}

	ReverseSlice(slice1)
	ReverseSlice(slice2)

	fmt.Println(slice1)
	fmt.Println(slice2)

}

func ReverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
