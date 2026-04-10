package main

import (
	"fmt"
)

func main() {
	// // Print the first 12 even numbers
	// count := 0
	// for i := 0; count < 12; i += 2 {
	// 	count++
	// 	fmt.Println(i)
	// }

	// // Print multiples of 5 from 5 to 50
	// for j := 5; j <= 50; j++ {
	// 	if j % 5 == 0 {
	// 		fmt.Println(j)
	// 	}
	// }

	// // Print a countdown from 10 to 1
	// for k := 10; k >= 1; k-- {
	// 	fmt.Println(k)
	// } 

	// // Print the numbers from 1 to 20, but skip multiples of 3
	// for l := 1; l <= 20; l++ {
	// 	if l % 3 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(l)
	// }

	// // Print the first 8 odd numbers using a counter
	// counter := 0
	// for m := 1; counter < 8; m += 2 {
	// 	counter++
	// 	fmt.Println(m)
	// }

	// // Sum an array
	// // Find max
	// // Reverse in place

	// Sum an array
	arr := []int{4, 8, -3, 10, 5}
	arr = []int{-7, 2, 9, 0, 6}
	arr = []int{1}
	sum := 0

	for n := 0; n < len(arr); n++ {
		sum += arr[n]
	}
	fmt.Println(sum)

	for _, n := range arr {
		sum += n
	}
	fmt.Println("Sum = ", sum)

	// Find max
	arr1 := []int{12, 4, 19, 8, 6}
	arr1 = []int{-10, -3, -50, -1}
	arr1 = []int{7}
	max := arr1[0]

	for _, p := range arr1 {
		if p >= max {
			max = p
		}
	}
	fmt.Println("Max =", max)

	// for p := 0; p < len(arr1); p++ {
	// 	if arr1[p] > max {
	// 		max = arr1[p]
	// 	}
	// }
	// fmt.Println("Max =", max)

	// Reverse in place
	arr2 := []int{1, 2, 3, 4, 5}
	arr2 = []int{10, 20, 30, 40}
	arr2 = []int{9, 8, 7}

	left := 0
	right := len(arr2) - 1

	for left < right {
		arr2[left], arr2[right] = arr2[right], arr2[left]

		left ++
		right --
	}
	fmt.Println("Reversed:", arr2)
}