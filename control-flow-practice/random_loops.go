// package main

// import "fmt"

// func main() {
// 	// Classic For-loop
// 	// Print all odd numbers from 1–100.
// 	max := 100
// 	for i := 1; i <= max; i++ {
// 		if i % 2 == 1 {
// 			fmt.Println("odd number", i)
// 		}
// 	}

// 	// While-loop style
// 	// print 1 to 100
// 	x := 1 // start from 1
// 	for x <= 100 {
// 		fmt.Println(x)
// 		x++
// 	}

// 	// Print all even numbers from 1–100.
// 	j := 1; // the actual digits
// 	for j <= 100 { // numbers less or 100
// 		if j % 2 == 0{
// 			fmt.Println("even number", j)
// 		}
// 		j++ // repitition
// 		// stops at 100
// 	}

// 	// sum 1 to 100
// 	sum := 0 // container
// 	y := 1 // the actual digits
// 	for y <= 100 { // numbers less or 100
// 		fmt.Printf("%d + %d = %d\n", y, sum, sum + y)
// 		sum += y // accumulation of the sum & numbers
// 		y++ // repitition
// 		// stops at 100
// 	}

// 	// Print all numbers 1–50, but skip multiples of 3.
// 	for i := 1; i <= 50; i++ {
// 		if i % 3 == 0 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}

// 	// Print numbers 1–20, stop if you reach a number divisible by 7.
// 	for i := 1; i <= 20; i++ {
// 		if i % 7 == 0 {
// 			break
// 		}
// 		fmt.Println(i)
// 	}

// 	// Print first 10 odd numbers using classic for loop.
// 	for i := 0; i <= 10; i++ {
// 		if i % 2 == 1 {
// 			fmt.Println(i)
// 		}
// 	}

// 	// Ask the user to enter numbers until they enter 0, sum all numbers.
// 	var numb int
// 	fmt.Println("Enter any number (enter zero to stop);")

// 	for numb >= 0 {
// 		fmt.Print("Number: ")
// 		fmt.Scan(&numb)
// 		if numb >= 0 {
// 			fmt.Println("You entered:", numb)
// 		}
// 	}
// 	fmt.Println("Stopped because you entered zero")

// 	// Print numbers 1–30, skip multiples of 2 and break at 25.
// 	for i := 1; i <= 30; i++ {
// 		if i % 2 == 0 {
// 			continue
// 		}
// 		if i == 25 {
// 			break
// 		}
// 	} 

// 	// task % and if-else
// /* 
// 	% means modulo and it gives the remainder
// 	if-else is to check the statement based on boolean condition

// 	I want to use this to generate remainder of simple arithmetic
// */
// 	num1 := 93
// 	// print all the odd numbers starting from 1

// 	// print the num1
// 	if num1 % 2 == 1 {
// 		fmt.Println("odd number", num1)
// 	} else {
// 		fmt.Println("even number", num1)
// 	}

// 	// Print all even numbers from 1–100.
// 	p := 1; // the actual digits

// 	for p <= 100 { // numbers less or 100
// 		if p % 2 == 0{
// 			fmt.Println("even number", p)
// 		}
// 		p++ // repitition
// 		// stops at 100
// 	}

// 	// Classic For-loop
// 	// Print all odd numbers from 1–100.
// 	maxi := 100
// 	for i := 1; i <= maxi; i++ {
// 		if i % 2 == 1 {
// 			fmt.Println("odd number", i)
// 		}
// 	}

// 	// While-loop style
// 	// print 1 to 100
// 	s := 1 // start from 1
// 	for x <= 100 {
// 		fmt.Println(s)
// 		s++
// 	}

// 	// Print all even numbers from 1–100.
// 	k := 1; // the actual digits
// 	for k <= 100 { // numbers less or 100
// 		if k % 2 == 0{
// 			fmt.Println("even number", k)
// 		}
// 		k++ // repitition
// 		// stops at 100
// 	}

// 	// sum 1 to 100
// 	addtn := 0 // container
// 	c := 1 // the actual digits
// 	for c <= 100 { // numbers less or 100
// 		fmt.Printf("%d + %d = %d\n", c, addtn, addtn + c)
// 		addtn += c // accumulation of the sum & numbers
// 		c++ // repitition
// 		// stops at 100
// 	}

// 	// Print all numbers 1–50, but skip multiples of 3.
// 	for i := 1; i <= 50; i++ {
// 		if i % 3 == 0 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}

// 	// Print numbers 1–20, stop if you reach a number divisible by 7.
// 	for i := 1; i <= 20; i++ {
// 		// if i / 7 == 2 {
// 		if i % 7 == 0 {
// 			break
// 		}
// 		fmt.Println(i)
// 	}

// 	// Print first 10 odd numbers using classic for loop.
// 	for i := 1; i <= 20; i+=2 { // skip even numbers
// 		if i % 2 == 1 {
// 			fmt.Println(i)
// 		}
// 	}

// 	count := 0
// 	for i := 1; count < 10; i += 2 {
// 		fmt.Println(i)
// 		count++
// 	}


// 	// Ask the user to enter numbers until they enter 0, sum all numbers.
// 	var numbr int // number to enter
// 	var summ int // sum container

// 	// get the input before starting the loop
// 	fmt.Println("Enter any number (enter zero to stop);")
// 	// First input BEFORE the loop
// 	fmt.Print("Number: ")
// 	fmt.Scan(&numbr)

// 	for numbr != 0 { // whatever number as long as it's not zero
// 		summ += numbr
// 		fmt.Println("The current sum", summ)

// 		// take the current input number
// 		fmt.Print("Number: ")
// 		fmt.Scan(&numbr)
// 	}

// 	fmt.Println("Stopped because you entered zero")
// 	fmt.Println("Finale sum: ", summ)

// // Print numbers 1–30, skip multiples of 2 and break at 25.
// 	for i := 1; i <= 30; i++ {
// 		if i % 2 == 0 {
// 			continue
// 		}
// 		fmt.Println(i)
// 		if i == 25 {
// 			break
// 		}
// 	}

// }