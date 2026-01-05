// package main

// import "fmt"

// func main() {
// 	arr := [4]int{3, 6, 9, 12}
// 	fmt.Println(arr)

// 	arr2 := arr
// 	fmt.Println(arr2)

// 	arr[1] = 4
// 	fmt.Println(arr)
// 	fmt.Println(arr2)

// 	arr2[3] = 8
// 	fmt.Println(arr2)
// 	fmt.Println(arr)

// 	arr4 := [3] int{9, 8, 2}
// 	fmt.Println(arr4)
// 	arr4[2] = 7
// 	fmt.Println(arr4)

// 	arr5 := [3] int{9, 8, 7}
// 	fmt.Println(arr4 == arr5)

// 	for _, v := range arr4 {
// 		fmt.Println(v)
// 	}

// 	// fixed array
// 	fixedArr := [...]string{"Oyindamola", "Abiola", "Susana", "Abosede"}
// 	fmt.Println(fixedArr)
// 	fmt.Println(len(fixedArr))

// 	carBrand := [...]string{"benz", "bmw", "toyota", "volvo", "corolla", "ford"}
// 	fmt.Println(carBrand[2])

// 	carBrand[2] = "mazda" 
// 	fmt.Println(carBrand)
// 	fmt.Println(len(carBrand))

// 	// Nested Array must be same length
// 	arrNested := [3][3]int{{2, 4}, {3, 6, 9}}
// 	fmt.Println(arrNested)


// 	// SLICES || DYNAMIC ARRAYS

// 	slc := []int{1993, 1774, 1634, 1842}
// 	fmt.Println(slc)
// 	slc = append(slc, 1679)
// 	fmt.Println(slc)

// 	slc2 := []string{"Abosede", "Susanna", "Dasola"}
// 	fmt.Println(len(slc2))

// 	// convert an array to a slice (create a new variable for the slice)
// 	fixedArr = [...]string{"Oyindamola", "Abiola", "Susana", "Abosede"}
// 	fmt.Println(fixedArr)
// 	fixedArrSlice := fixedArr[:]
// 	fixedArrSlice = append(fixedArrSlice, "Efunroye")
// 	fmt.Println(fixedArrSlice)
// }

// 		// SLICE: make & append
// 	// slice are generally and by default flexible
// func main() {
// 	// slice1 - fixed but expandable
// 	fixedSlice := make([]int, 5, 8) // original 5items [0,0,0,0,0]
// 	fmt.Println(len(fixedSlice), cap(fixedSlice)) // 5, 8

// 	fixedSlice = append(fixedSlice, 6, 8) // added 2 items
// 	fmt.Println(len(fixedSlice), cap(fixedSlice)) // 7, 8

// 	// appending it makes it create a copy of fixedSlice and expand the copy
// 	fixedSlice = append(fixedSlice, 1, 2, 3, 4) // appended 4 more items [0, 0, 0, 0, 0, 1, 2, 3, 4]
// 	fmt.Println(len(fixedSlice), cap(fixedSlice)) // 11, 16

// 	// slice2 - dynamic slice
// 	dynamicSlice := [6]int {}
// 	fmt.Println(len(dynamicSlice), cap(dynamicSlice))
// }