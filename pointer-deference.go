package main

import "fmt"

// func updateName(lastName *string) {
// 	*lastName = "Abiola"
// 	fmt.Println(lastName)
// }

// func main() {
// 	name := "Oyindamola"
// 	pointer := &name
// 	fmt.Println("The pointer is: ", pointer) // gets the memory of the variable
// 	fmt.Println("The pointer is: ", *pointer) // dereference gets the value

// 	oldLastName := "Rasak"
// 	oln := &oldLastName // this is the variable location
// 	updateName(oln)
// 	fmt.Println(oldLastName)

// }

func updateAges(ages *[]int) []int {
	for i := range *ages {
		(*ages)[i] *= 2
	}
	return *ages
}

func main() {
	oldAges := []int{22, 24, 26, 28, 30}
	updateAges(&oldAges)

	fmt.Println(oldAges)
}

