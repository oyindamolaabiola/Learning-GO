// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// /** 
// 	Sum positive numbers until user enters a negative number
// 	Keep adding positive numbers.
// 	Stop the loop when the user enters any negative number.
// 	Print the final sum.
// **/
// 	var num int
// 	var sum int

// 	// ask the user for number before the loop logic
// 	fmt.Println("Enter any number (enter negative number to stop);")
// 	fmt.Print("Number: ")
// 	fmt.Scan(&num)

// 	for num >= 0 {
// 		sum += num

// 		fmt.Println("Current sum: ", sum)
// 		fmt.Print("Number: ")
// 		fmt.Scan(&num)
// 	}

// 	fmt.Printf("Stopped because you entered a negative number")
// 	fmt.Println("Finale sum: ", sum)

// // /** 
// // 	Count how many even numbers the user enters
// // 	Ask the user for numbers repeatedly.
// // 	Stop when the user enters 0.
// // 	Count only the even numbers (2, 4, 6…).
// // **/
// 	var numb int
// 	var evenNumbs int

// 	// ask the user for number before the loop logic
// 	fmt.Println("Enter any number (enter zero to stop);")
// 	fmt.Print("Number: ")
// 	fmt.Scan(&numb)

// 	for numb != 0 {
// 		if numb % 2 == 0 {
// 			evenNumbs++
// 		}

// 		fmt.Println("Current count of even numbers: ", evenNumbs)

// 		fmt.Print("Number: ")
// 		fmt.Scan(&numb)
// 	}

// 	fmt.Println("Stopped because you entered zero")
// 	fmt.Println("Finale count: ", evenNumbs)

// // /**
// // 	Ask the user for names until they type “stop”
// // 	Keep reading a string.
// // 	Stop only when the user types "stop" (case-insensitive).
// // **/
// 	var name string
// 	var allNames []string

// 	// ask user before the loop logic
// 	fmt.Println("Enter your names, (enter 'stop' to stop)")
// 	fmt.Print("Name: ")
// 	fmt.Scan(&name)

// 	// for name != "stop" { ??
// 	for strings.ToLower(name) != "stop" {
// 		allNames = append(allNames, name) 
// 		fmt.Println("Current name list: ", allNames)
// 		fmt.Print("Name: ")
// 		fmt.Scan(&name)
// 	}

// 	fmt.Println("Stopped because you entered 'stop'")
// 	fmt.Println("Finale name list: ", allNames)

// // /**
// // 	Find the largest number the user enters
// // 	Keep reading numbers.
// // 	Stop on 0.
// // 	Track the maximum number entered. 
// // **/
// 	var number int
// 	var largest int

// 	fmt.Println("Enter any number (enter zero to stop);")
// 	fmt.Print("Number: ")
// 	fmt.Scan(&number)

// 	largest = number // assume first number is largest (if it's zero, loop will exit)

// 	for number != 0 {
// 		if number > largest {
// 			largest = number
// 		}

// 		fmt.Print("Number: ")
// 		fmt.Scan(&number)
// 	}

// 	fmt.Println("Stopped because you entered zero")
// 	fmt.Println("Largest number:", largest)


// /**	
// 	Ask the user for numbers until they enter 0, but ignore negatives
// 	Only sum positive numbers.
// 	Stop when user enters 0.
// 	Do nothing when user enters a negative number.
// **/
// 	var numbr int
// 	var summ int

// 	// ask the user for number before the loop logic
// 	fmt.Println("Enter any number (enter zero to stop);")
// 	fmt.Print("Number: ")
// 	fmt.Scan(&numbr)

// 	for numbr != 0 {
// 		if numbr > 0 {
// 			summ += numbr
// 		}
// 		fmt.Println("Current number: ", numbr)

// 		fmt.Print("Number: ")
// 		fmt.Scan(&numbr)
// 	}

// 	fmt.Println("Stopped because you entered zero")
// 	fmt.Println("Sum of the positive numbers: ", summ)
// }