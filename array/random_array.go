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
	// arr6 := []int{1, 2, 3, 5, 4, 6}
	arr6 := []int{10, 20, 15, 30}
	sorted := true // assume true until proven otherwise

	// start from index 1 not zero so that i-1 will be within the range
	for i := 1; i < len(arr6); i++ {
		if arr6[i] < arr6[i-1] {
			sorted = false
			break
		}
	}

	fmt.Println(sorted)

	// for-range
	for ind := range arr6 { // ignore values, not index
		if ind == 0 {
			continue
		}
		if arr6[ind] < arr6[ind-1] {
			sorted = false
			break
		}
	}

	fmt.Println(sorted)

	// Find the second‑largest element (invariant)
	arr7 := []int{12, 7, 19, 3, 19, 5}

	// assume the indices of max and second max
	maxx := arr7[0]
	secMax := arr7 [1]

	for i := 0; i < len(arr7); i++ {
		if arr7[i] > maxx {
			secMax = maxx
			maxx = arr7[i]
		} else if arr7[i] < maxx && arr7[i] > secMax {
			secMax = arr7[i]
		}
	}
	fmt.Println("Second Largest: ", secMax)

	// Deduplicate an array
	arr8 := []int{4, 2, 4, 6, 2, 8, 6}
	arrMap := make(map[int]bool) // for validation
	var newArr8 []int

	for _, m := range arr8 {
		if !arrMap[m] {
			newArr8 = append(newArr8, m)
			arrMap[m] = true
		} 
	}
	fmt.Println("New Array: ", newArr8)

	// Check if a Slice Is Non-Decreasing
	arr9 := []int{2, 4, 4, 7, 9}
	// assume the data is non-decreasing (the curr > or = prev)
	nonDec := true

	for n := 1; n < len(arr9); n++ {
		if arr9[n] < arr9[n-1] {
			nonDec = false
			break
		}
	}
	fmt.Println(nonDec)

	// Check if a Slice Is strictly increasing
	arr10 := []int{2, 4, 4, 7, 9}
	// assume the array is strictly increasing (curr > prev)
	strInc := true

	for o := 1; o < len(arr10); o++ {
		if arr10[o] <= arr10[o-1] {
			strInc = false
			break
		}
	}
	fmt.Println(strInc)

	// // Find the smallest and second smallest number
	// arr := []int{8, 3, 5, 1, 9, 2}

	// // Find the largest odd number
	// arr := []int{12, 7, 19, 3, 19, 5}

	// // Find the second largest DISTINCT number
	// arr := []int{10, 20, 20, 30, 30, 25}

}