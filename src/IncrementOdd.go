package main

import "fmt"

func main() {

	// result#1: 1 3 3 5 5 7 7 9 9 11
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//To check, I specifically wrote: the number 3 is located in even places, and the number 1 is in odd places.
	// result#2: 3 2 3 2 3 2 3 2 3 2 3 2 3
	slice2 := []int{3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3}

	//Randomly
	// result#3: 12 21 0 6 6 799 81 55 36 20 29 48 144 8
	slice3 := []int{12, 20, 0, 5, 6, 798, 81, 54, 36, 19, 29, 47, 144, 7}

	IncrementOdd(slice1)
	IncrementOdd(slice2)
	IncrementOdd(slice3)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

func IncrementOdd(slice []int) {
	for i := 0; i < len(slice); i++ {
		if i%2 != 0 {
			slice[i]++
		}
	}
}
