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

	// // Reverse the array in place – Show the slice after reversal.
	// arr4 := []int{1, 2, 3, 4, 5} 

	// for _, l := range arr4 {
	// 	arr4[4] = arr4[0]
	// 	arr4[3] = arr4[1]
	// }
	// fmt.Println("Slice = ", arr4)

	// Remove all negatives (new slice)
	arr5 := []int{-3, 5, -1, 0, 8, -2, 7}
	var negArr5 []int

	for _, j := range arr5 {
		if j < 0 {
			negArr5 = append(negArr5, j)
		}
	}
	fmt.Println("Negative = ", negArr5)
}