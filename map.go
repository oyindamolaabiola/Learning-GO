package main

import "fmt"

func main() {

	// teacher := map[string]string {
	// 	"name": "odunayo",
	// 	"age": "25",
	// 	"country": "nigeria",
	// 	"hobby": "sleeping",
	// }

	// Print out the map
	// fmt.Println(teacher) // print raw data
	// // 2.
	// for key, value := range teacher {
	// 	fmt.Printf("%v: %v\n", key, value)
	// }

	// Update map value using the key 
	// teacher["name"] = "Oladunni"
	// fmt.Println(teacher["name"])

	// Update map using loop
	// for key, value := range teacher {
	// 	if key == "hobby" {
	// 		teacher[key] = "studying"
	// 	}
	// 	// add property
	// 	teacher["complexion"] = "Choco-Brown"
	// 	fmt.Printf("%v: %v\n", key, value)
	// }

	// // Add properties
	// teacher["state"] = "oyo-state"
	// teacher["weight"] = "45"

	// // else-if in forloop
	// present := false
	// for key := range teacher {
	// 	if key == "height" {
	// 		fmt.Println("Height found")
	// 		present = true
	// 	} else if key == "weight" {
	// 		fmt.Println("Weight is", teacher["weight"])
	// 		present = true
	// 	}
	// }

	// if !present {
	// 	fmt.Println("Non is present")
	// }

	atheletes := map[string]int {
		"nigeria": 234,
		"ghana": 233,
		"kenya": 254,
		"zambia": 260,
		"tanzania": 255,
	}

	// fmt.Println(atheletes)
	// for key, value := range atheletes {
	// 	fmt.Printf("%v: %v\n", key, value)
	// }

	// // Prints the values only
	// for key := range atheletes {
	// 	fmt.Println(atheletes[key])
	// }

	// ADD PROPERTY
	atheletes["uganda"] = 256
	atheletes["cote d'ivoire"] = 225
	// fmt.Println(atheletes["uganda"])

	// for key, value := range atheletes {
	// 	fmt.Printf("Country: %v, code: %v\n", key, value)
	// }

	// Updating the value
	// for key := range atheletes {
	// 	if key == "ghana" {
	// 		atheletes[key] = 222
	// 		fmt.Println(atheletes["ghana"])
	// 	}
	// 	// fmt.Println(atheletes["ghana"])
	// }

	// // Updating to key
	// for key, value := range atheletes {
	// 	if value == 260 {
	// 		key = "Benin"
	// 	}
	// 	fmt.Printf("%v: %v\n", key, value)
	// }

	// Accessing the map properties
	// fmt.Println(atheletes["cote d'ivoire"]) // access value

	// for key, value := range atheletes { // access the key by loop range
	// 	if value == 234 {
	// 		fmt.Println(key)
	// 	}
	// }

	if atheletes["nigeria"] == 234 {
		fmt.Println("Nigeria's code is 234")
	} else {
		fmt.Println("Nigeria is not present")
	}

	available := false
	for value := range atheletes {
		if value == "zambiia" {
			fmt.Println("yes, zambia is 261")
			available = true
			break
		} else if value == "tanzaniia" {
			fmt.Println("yes, tanzania is 255")
			available = true
			break
		}
	}

	if !available {
		{
			fmt.Println(atheletes["zambia"])
		}
	}

	// Create map
	// states := make(map[string]string)
	// states["lagos"] = "ikeja"
	// states["oyo"] = "ibadan"
	// states["abuja"] = "fct"
	// states["ogun"] = "abeokuta"
	// fmt.Printf("%v\n", states)

	// Accessing the map properties
	// fmt.Println(states["ogun"])

	// found := false
	// for _, value := range states {
	// 	if value == "kaduna" {
	// 		fmt.Println("Kaduna is present")
	// 		found = true
	// 	} else if value == "abeok" {
	// 		fmt.Println("Got Abeokuta here:", value)
	// 		found = true
	// 	}
	// }

	// if !found {
	// 	for key := range states {
	// 		fmt.Println(key)
	// 	}
	// }
	

	// Create map 2
	// student := make(map[string]int)
	// student["marvelous"] = 2001
	// student["faith"] = 2004
	// student["obaloluwa"] = 2007
	// student["blessing"] = 2009
	// // fmt.Printf("%v\n", student)

	// // Accessing the map properties
	// fmt.Println(student["faith"])
	// for key, value := range student {
	// 	if value == 2001 {
	// 		fmt.Println(key)
	// 	}
	// }

	// OTHER TYPES
	// game := map[int]bool {
	// 	2: true,
	// 	1: false,
	// 	4: true,
	// 	3: false,
	// }
	// fmt.Printf("%v\n", game)

	// game2 := map[string]bool {
	// 	"lovely": true,
	// 	"toxic": false,
	// 	"compassionate": true,
	// 	"mean": false,
	// }
	// fmt.Printf("%v\n", game2)

	// height := map[string]float32 {
	// 	"bisola": 5.2,
	// 	"feyisara": 5.2,
	// 	"akindewa": 5.0,
	// 	"oyindamola": 5.4,
	// }
	// fmt.Printf("%v\n", height)

	// Accessing the map properties

	// convert an array to a map
	// ADD KEY <=> to the dictionary

	// // Print all the map
	// fmt.Printf("%v\n", student)
	// 
	// fmt.Printf()

}
