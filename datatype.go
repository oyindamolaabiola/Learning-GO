package main

import "fmt"

// POINTER
// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	ptr := &arr[2]
// 	fmt.Println(*ptr)

// 	*ptr = 10 // allocate 10 to storage 2
// 	// print arr at index 2
// 	fmt.Println(arr[2])
// }

// GENERIC
// This was not allowed befor G0 1.18
// func Testgeneric[X any](value X) {
// 	fmt.Println(value)
// }

// func main() {
// 	Testgeneric(23)
// 	Testgeneric("August")
// }

/* VAVRIABLE DECLARATION */
func main() {
	// variable_list type of declaration
	// var i, j, k int
	// var a, b float

	// i = 2, j = 19, k = 5

	// static variable declaration
	// var a int
	// a = 23
	// fmt.Println(a)
	// fmt.Printf("a is of type %T\n", a)

	// // dynamic variable declaration
	// var m int = 32
	// n := 2.0

	// fmt.Println(m)
	// fmt.Println(n)
	// fmt.Printf("m is a type %T\n", m)
	// fmt.Printf("n is a type %T\n", n)

	// fmt.Println("Oyindamola\tDasola!")
	// fmt.Println("Oyindamola\bDasola!")
	// // fmt.Println("Oyindamola\nDasola!")
	// fmt.Println("Oyindamola\rDasola!")
	// const LENGTH int = 20
	// const width int = 10

	// a--
	// fmt.Printf("Line 6 - Value of a is %d\n", a)
	// a++
	// fmt.Printf("Line 7 - Value of a is %d\n", a)

	// var k int = 21
	// var l int = 40

	// if k <= l {
	// 	fmt.Printf("Line 4 - k is either less than or equal to l\n")
	// }
	// if l >= k {
	// 	fmt.Printf("Line 5 - b is either greater than or equal to l\n")
	// }

	// var a uint = 60 /* 60 = 0011 1100 */
	// var b uint = 13 /* 13 = 0000 1101 */
	// var c uint = 0
	// c = a & b /* 12 = 0000 1100 */
	// fmt.Printf("Line 1 - Value of c is %d\n", c)
	// c = a | b /* 61 = 0011 1101 */
	// fmt.Printf("Line 2 - Value of c is %d\n", c)
	// c = a ^ b /* 49 = 0011 0001 */
	// fmt.Printf("Line 3 - Value of c is %d\n", c)

	// c = a << 2 /* 240 = 1111 0000 */
	// fmt.Printf("Line 4 - Value of c is %d\n", c)
	// c = a >> 2 /* 15 = 0000 1111 */
	// fmt.Printf("Line 5 - Value of c is %d\n", c)

	// SCOPE
	// variable can be declared outside of the function in GO meaning a global variable also called package-level variable

		/* VARIABLES */
	// Explicit variables
	// var name string = "Ayotolani"
	// var year int = 1996
	// var height float64 = 5.26
	// var isPretty bool = true

	// // Explicit Inline multi-variable but both has to be the same var-type
	// var firstName, lastName string = "Odunayo", "Abidoye"

	// // Inferred variables
	// name2 := "Bisola"
	// year2 := 1995
	// height2 := 5.3
	// isNotPretty := false

	// // Inferred works as Inline multi-variables too
	// state, num := "Lagos", 12

	/* COMPARISON */
	// Strings
	// isString := (name < name2)
	// isEqualAnd := !!isString
	// fmt.Printf("Are both the same: %t\n", isEqualAnd)

	// Boolean
	// isBoolean := (isPretty && isNotPretty)
	// isItReallyTrue := !!isBoolean
	// fmt.Printf("Are both true/false: %t\n", isBoolean)
	// fmt.Printf("Is it really true/false: %t\n", isItReallyTrue)

	// Float
	// isHeight := (height < height2)
	// fmt.Printf("Are both floats the same? %t\n", isHeight)


	/* CONTROL STRUCTURES */
	// For loop
	for i := 0; i < 10; i++ {
		if (i % 2 == 1) {
			fmt.Println((i))
		}
	}

	// // While loop ()
	// i := 0;
	// for i <= 5 {
	// 	fmt.Println((i))
	// 	i++
	// } 

	// Range // Array
	// num := []int{1, 3, 5, 7, 9} // dynamic
	// ages := [5]int{}
	// for i, value := range num {
	// 	fmt.Printf("Index: %d, Value: %d\n", i, value)
	// }

	// Array String
	// names := []string{"Abisola", "Feyintola", "Ayotolani", "Mariam", "Asake"}
	// for i, name := range names {
	// 	fmt.Printf("Name %d: %s\n", i, name)
	// }
	
}

