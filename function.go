package main

import (
	// "errors"
	"fmt"
)

func greet(name string) {
	fmt.Println("Hello", name)
}

func familyName(firstName string) {
	fmt.Println("Hello", firstName, "Abiola")
}

func add(a int, b int) int {
	return a + b 
}

func sub(c int, d int) int {
	return c - d
}

// func div(e, f int) (int, error) {
// 	if f == 0 {
// 		return 0, errors.New("Division by zero error") 
// 	}
// 	return e / f, nil
// }

// func main() {
// 	greet("Oyindamola")
// 	familyName("Oyindamola")
// 	familyName("Dasola")
// 	familyName("Susanna")
// 	familyName("Efunroye")
// 	addResult := add(2, 4)
// 	subResult := sub(16, 7)
// 	divResult, err := div(24, 8)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Result:", divResult)
// 	}


// 	fmt.Println(addResult)
// 	fmt.Println(subResult)
// 	// fmt.Println(divResult)
// }


// func mod(div1, div2 int) (int, error) {
// 	ans := div1 / div2

// 	if ans == 1 {
// 		return 0, errors.New("Div1 is not equally divisible by div2")
// 	}
// 	return ans, nil
// }

// func main() {
// 	mainAns, err := mod(100, 30)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Output:", mainAns)
// 	}

// }


func modifyArr(arr [5]int) {
	for i := range arr {
		arr[i] *= 2
	}

	fmt.Println("Inside modifyArr:", arr)
}

func main() {
	// rslt := modifyArr(2, 4, 6, 8, 10)

	arr2 := [5] int{2, 4, 6, 8, 10}
	modifyArr(arr2)
	// fmt.Println("After modifyArr:", arr2)
}