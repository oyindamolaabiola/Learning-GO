// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	// Sum until negative — Read integers; stop on any negative; print sum.

// 	var sum int
// 	var intgr int
	
// 	fmt.Println("Enter any number (enter any negative to stop);")
// 	fmt.Print("Number: ")
// 	fmt.Scan(&intgr)

// 	for intgr >= 0 {
// 		sum += intgr
// 		fmt.Println("Current sum is: ", sum)

// 		fmt.Print("Number: ")
// 		fmt.Scan(&intgr)
// 	}

// 	fmt.Println("Stopped because you entered a negative integer")
// 	fmt.Println("Total sum:", sum)

// 	// Count evens until 0 — Count even inputs; stop when 0 is entered.
// 	var digit int
// 	var evenNumbs int
	
// 	fmt.Println("Enter any number (enter zero to stop);")
// 	fmt.Print("Number: ")
// 	fmt.Scan(&digit)

// 	for digit != 0 {
// 		if digit % 2 == 0 {
// 			evenNumbs ++
// 		}
// 		fmt.Println("Current even number count: ", evenNumbs)

// 		fmt.Print("Number: ")
// 		fmt.Scan(&digit)
// 	}

// 	fmt.Println("Stopped because you entered zero")
// 	fmt.Println("Total even number count:", evenNumbs)

// 	// Collect names until "stop" — Append strings to slice; stop case-insensitive at "stop".
// 	var name string
// 	var nameList []string
	
// 	fmt.Println("Enter any number (enter any negative to stop);")
// 	fmt.Print("Name: ")
// 	fmt.Scan(&name)

// 	for name != strings.ToLower("stop") { // "Stop" didn't work, STOP & stop worked
// 		nameList = append(nameList, name)
// 		fmt.Println("Current name list: ", nameList)

// 		fmt.Print("Name: ")
// 		fmt.Scan(&name)
// 	}

// 	fmt.Println("Stopped because you entered 'stop'.")
// 	fmt.Println("Complete name list:", nameList)

// 	// Average until 0 — Read numbers until 0; print average (watch division by zero).
// 	var num int
// 	var average int
// 	var numbCount int
	
// 	fmt.Println("Enter zero to stop;")
// 	fmt.Print("Number: ")
// 	fmt.Scan(&num)

// 	for num != 0 {
// 		numbCount ++
// 		average += num
// 		// average = average / numbCount
// 		fmt.Println("Current average: ", average)

// 		fmt.Print("Number: ")
// 		fmt.Scan(&num)
// 	}
// 	average = average / numbCount
// 	fmt.Println("Stopped because you entered zero")
// 	fmt.Println("Total average:", average)

// 		// Sum positives only until 0 — Ignore negatives; stop at 0; print sum.
// 		var numbr int
// 		var totalPostv int
		
// 		fmt.Println("Enter zero to stop;")
// 		fmt.Print("Number: ")
// 		fmt.Scan(&numbr)

// 		for numbr != 0 {
// 			if numbr > 0 {
// 				totalPostv += numbr
// 			}
// 			fmt.Println("Sum of the positive numbers: ", totalPostv)

// 			fmt.Print("Number: ")
// 			fmt.Scan(&numbr)
// 		}

// 		fmt.Println("Stopped because you entered zero")
// 		fmt.Println("Sum of the positives:", totalPostv)

// 	// Find max until 0 — Track the largest number; stop when 0 entered. (Prime with first input.)
// 	var numm int
// 	var largestNumb int

// 	fmt.Println("Enter zero to stop;")
// 	fmt.Print("Number: ")
// 	fmt.Scan(&numm)

// 	largestNumb = numm

// 	for numm != 0 {
// 		if numm > largestNumb {
// 			largestNumb = numm
// 		}
// 	fmt.Println("Current largest number: ", largestNumb)	

// 	fmt.Print("Number: ")
// 	fmt.Scan(&numm)
// 	}

// 	fmt.Println("Stopped because you entered zero")
// 	fmt.Println("The largest number: ", largestNumb)

// 	// Find min until 0 — Track smallest number; stop on 0.
// 	var nummb int
// 	var minNumb int

// 	fmt.Println("Enter zero to stop;")
// 	fmt.Print("Number: ")
// 	fmt.Scan(&nummb)

// 	minNumb = nummb

// 	for nummb != 0 {
// 		if nummb < minNumb {
// 			minNumb = nummb
// 		}
// 		fmt.Println("Current min number: ", minNumb)	

// 		fmt.Print("Number: ")
// 		fmt.Scan(&nummb)
// 	}

// 	fmt.Println("Stopped because you entered zero")
// 	fmt.Println("The min number: ", minNumb)

// 	// Count odd numbers until multiple of 10 — Stop when a number divisible by 10 is entered.
// 	var numbbr int
// 	var oddCount int 

// 	fmt.Println("Enter multiple of 10 to stop;")
// 	fmt.Print("Number: ")
// 	fmt.Scan(&numbbr)

// 	for numbbr >= 0 {
// 		if numbbr % 2 == 1 {
// 			oddCount++
// 		} else if numbbr % 10 == 0 {
// 			break
// 		}
// 		fmt.Println("Current odd number count: ", oddCount)	

// 		fmt.Print("Number: ")
// 		fmt.Scan(&numbbr)
// 	}

// 	fmt.Println("Stopped because you entered multiple of 10")
// 	fmt.Println("The final odd number count: ", oddCount)

// 	// Skip range — Sum inputs until -1, but ignore any number in range [20,30].
// 	var n int
// 	var numSum int 

// 	fmt.Println("Enter any number but -1 to stop;")
// 	fmt.Print("Number: ")
// 	fmt.Scan(&n)

// 	for n >= 0 {
// 		if n >= 20 && n <=30 {
// 			// ask for the next number to check the range
// 			fmt.Print("Number: ")
//         	fmt.Scan(&n)
// 			continue
// 		}
// 		numSum += n	
// 		fmt.Println("Current sum is: ", numSum)
		
// 		fmt.Print("Number: ")
// 		fmt.Scan(&n)
// 	}

// 	fmt.Println("Stopped because you entered a negative number")
// 	fmt.Println("The final sum: ", numSum)
// 	// what if the user enters -2 or -3??

// 	// Password prompt — Keep asking for string until it equals a password (exact match).
// 	word := "Oyindasola"
// 	var password string

// 	fmt.Println("Enter the correct password: ")
// 	fmt.Print("Password: ")
// 	fmt.Scan(&password)

// 	for password != word {
// 		fmt.Println("Wrong password!")

// 		fmt.Print("Password: ")
// 		fmt.Scan(&password)
// 	}
// 	fmt.Println("Yay!, password is a match: ", password)

// 	// First 12 evens — Use classic for (no if) to print first 12 even numbers.
// 	var countrr int
// 	for i := 0; countrr < 12; i += 2 {
// 		fmt.Println(i)
// 		countrr++
// 	}

// 	// Multiples of 5 — Print multiples of 5 from 5 to 50 inclusive.
// 	for j := 5; j <= 50; j += 5 {
// 		fmt.Println(j)
// 	}

// 	// Countdown — Print 10 down to 1 using a decrementing for.
// 	for k := 10; k >= 1; k-- {
// 		fmt.Println(k)
// 	}

// 	// 1–20 skip multiples of 3 — Use continue inside a classic for.
// 	for l := 1; l <=20; l++ {
// 		if l % 3 == 0 {
// 			continue
// 		}
// 		fmt.Println(l)
// 	}

// 	// First 8 odds using counter — i := 1 with count external pattern.
// 	var countee int
// 	for m := 1; countee < 8; m += 2 {
// 		fmt.Println(m)
// 		countee++
// 	}

// 	// Running median (approx) — Read N numbers (first read N), compute 
// 	// running average after each input. (counted loop + accumulator)
// 	var p int
// 	var midCount int
// 	var median int

// 	fmt.Println("Enter 0 to stop;")
// 	fmt.Print("Number: ")
// 	fmt.Scan(&p)

// 	for p != 0 {
// 		midCount++
// 		median += p
// 		median = median / midCount
// 		fmt.Println("Current median: ", median)

// 		fmt.Print("Number: ")
// 		fmt.Scan(&p)
// 	}
// 	fmt.Println("Stopped because you entered zero")
// 	fmt.Println("Total average:", median)

// 	// Top 3 largest — Read until 0, maintain top 3 largest numbers seen (comparison pattern & classic shifting).
// 	var r int
// 	larg := [3] int{}

// 	fmt.Println("Enter 0 to stop;")

// 	fmt.Print("Number: ")
// 	fmt.Scan(&r)

// 	for r != 0 {
// 		// shift is needed first to fill the indexes
// 		if r > larg[0] {
// 			larg[2] = larg[1]
// 			larg[1] = larg[0]
// 			larg[0] = r
// 		} else if r > larg[1] {
// 			larg[1] = r
// 		} else if r > larg[2] {
// 			larg[2] = r
// 		}
// 		fmt.Println("Current list: ", larg)

// 		fmt.Print("Number: ")
// 		fmt.Scan(&r)
// 	}
// 	fmt.Println("Stopped because you entered zero")
// 	fmt.Println("Final largest 3 numbers: ", larg)

// 	// Unique names collector — Read names until "done", 
// 	// append only new names (use map for uniqueness).
// 	var names []string
// 	var namee string
// 	// initialize map
// 	namesMap := make(map[string]bool)

// 	fmt.Println("Enter done to stop;")
// 	fmt.Print("Name: ")
// 	fmt.Scan(&namee)

// 	for namee != strings.ToLower("done") {
// 		if !namesMap[namee] { // if name is not in the map or not true
// 			names = append(names, namee)
// 			namesMap[namee] = true // set as true
// 		}
// 		fmt.Println("Current name list: ", names) // print name list

// 		fmt.Print("Name: ")
// 		fmt.Scan(&namee)
// 	}
// 	fmt.Println("Stopped because you entered done")
// 	fmt.Println("Final name list: ", names)

// 	// Matrix-like loops — Print multiplication table 1–10 (nested for loops).
// 	var ans int
// 	for t := 1; t <= 10; t++ { // row-vertical-outer
// 		// column loop below complete-first before moving to the next row
// 		for u := 1; u <= 10; u++ { // column-horizontal-inner
// 			ans = t * u
// 			// fmt.Println(ans)
// 			fmt.Printf("%d x %d = %d \n", t, u, ans)
// 		}
// 	}

// 	// FizzBuzz variant — For 1..100, print fizz/buzz/fizzbuzz using classic for and conditionals.
// 	// overlapping & conflicts variant
// 	for v := 1; v <= 100; v++ {
// 		if v % 3 == 0 && v % 5 == 0 { // this runs first (logical ordering)
// 			fmt.Println("fizzbuzz")
// 		} else if v % 3 == 0 {
// 			fmt.Println("fizz")
// 		} else if v % 5 == 0 {
// 			fmt.Println("buzz")
// 		} else {
// 			fmt.Println(v)
// 		}
// 	}
// }