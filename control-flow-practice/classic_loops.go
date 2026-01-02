package main

import (
	"fmt"
)

func main() {
	// Print the first 12 even numbers
	count := 0
	for i := 0; count < 12; i += 2 {
		count++
		fmt.Println(i)
	}

	// Print multiples of 5 from 5 to 50
	for j := 5; j <= 50; j++ {
		if j % 5 == 0 {
			fmt.Println(j)
		}
	}

	// Print a countdown from 10 to 1
	for k := 10; k >= 1; k-- {
		fmt.Println(k)
	} 

	// Print the numbers from 1 to 20, but skip multiples of 3
	for l := 1; l <= 20; l++ {
		if l % 3 == 0 {
			continue
		}
		fmt.Println(l)
	}

	// Print the first 8 odd numbers using a counter
	counter := 0
	for m := 1; counter < 8; m += 2 {
		counter++
		fmt.Println(m)
	}
}