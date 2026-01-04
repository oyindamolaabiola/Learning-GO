package main

import (
	"fmt"
)

func main() {
	
	// Print a 4 × 5 grid of *
	for a := 1; a <= 4; a++ {
		for b := 1; b <= 5; b++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	// Print right angle triangle
	for row := 1; row <= 4; row++ {
		for star := 1; star <= row; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Print inverted right angle triangle
	for row := 1; row < 5; row++ {
		for star := 1; star <= 5-row; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Print a pyramid of height 4
	for row := 1; row < 5; row++ {
	// create the empty space to be inverted
		for space := 1; space <= 5-row; space++ {
			fmt.Print(" ")
		}
	// create the stars to be right angle triangle (2times)
		for star := 1; star <= row; star++ {
				fmt.Print("*")
		}

		for star := 1; star < row; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Print a 4×4 checkerboard using X and O.
	// Conditions based on (row + col) % 2
	for row := 1; row <= 4; row++ {
		for col := 1; col <= 4; col++ {
			if (row + col) % 2 == 0 {
				fmt.Print("X")
			} else {
				fmt.Print(0)
			}
		}
		fmt.Println()
	}

}
