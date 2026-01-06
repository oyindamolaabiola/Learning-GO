package main

import (
	"fmt"
)

/* 
	Before coding, ask:
	Am I visiting all elements or stopping early?
	What state do I track? (sum, max, count, map, slice)
	Am I keeping, skipping, or ranking elements?
	Does order matter?
*/

func main () {
	// Sum all elements in an array
	// : all elements : sum : none : nope
	arr1 :=	[]int{3, 7, -2, 5, 10}
	sum := 0

	for _, i := range arr1 {
		sum += i
	}
	fmt.Println("Sum = ", sum)

	// Find max and min – Return both values as a pair (max, min)
	arr2 := []int{12, 4, 19, 8, 6}
	max := arr2[0]
	min := arr2[0]
	var pair [2]int

	for _, j := range arr2 {
		if j >= max {
			max = j
		}
		pair[0] = max
		if j <= min {
			min = j
		}
		pair[1] = min
	}
	fmt.Println("Pair = ", pair)

	// Count even numbers – How many even integers are there?
	arr3 := []int{2, 9, 4, 7, 12, 15, 20}
	var evenNumber int

	for _, k := range arr3 {
		if k % 2 == 0 {
			evenNumber++
		}
	}
	fmt.Println("Even number = ", evenNumber)

	// Reverse the array in place – Show the slice after reversal.
	arr4 := []int{1, 2, 3, 4, 5} 

	// pointers
	left := 0 // first item
	right := len(arr4) - 1 // last item

	for left < right {
		// swap inwardly
		arr4[left], arr4[right] = arr4[right], arr4[left] 
		left++ // moves right
		right-- // moves left
	}
	fmt.Println("Reverse version: ", arr4)

	// Remove all negatives (new slice)
	arr5 := []int{-3, 5, -1, 0, 8, -2, 7}
	var nonNegArr5 []int

	for _, j := range arr5 {
		if j >= 0 {
			nonNegArr5 = append(nonNegArr5, j)
		}
	}
	fmt.Println("Non-negative = ", nonNegArr5)

	// Check if the array is sorted (ascending) – Return true if sorted, otherwise false.
	arr6 := []int{1, 2, 3, 5, 4, 6}
	

}