package main

import "fmt"

func main() {

	firstFunc := func(slice []int) {
		fmt.Println("firstFunc:", slice)
	}

	secondFunc := appendFunc(firstFunc, printFunc, incrementFunc)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10}

	secondFunc(slice1)
	secondFunc(slice2)
}

func appendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(slice []int) {

		dst(slice)

		for _, f := range src {
			f(slice)
		}
	}
}

func printFunc(slice []int) {
	fmt.Println("PrintFunc:", slice)
}

func incrementFunc(slice []int) {

	for i := range slice {
		slice[i]++
	}

}
