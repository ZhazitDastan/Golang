package main

import "fmt"

func main() {

	// result#1: 0 1 2 3 4 5 6 7 8 9 10
	slice1 := []int{10, 9, 8, 7, 6, 5, 3, 4, 1, 0, 2}

	// result#2: 11 21 89 128 505 654 987
	slice2 := []int{11, 21, 987, 128, 505, 654, 89}

	// result#3: 10 12 12 12 13 28 50 58 67 155
	slice3 := []int{12, 12, 10, 12, 58, 67, 13, 155, 28, 50}

	SortSlice(slice1)
	SortSlice(slice2)
	SortSlice(slice3)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

func SortSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
