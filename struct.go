package main

import "fmt"

type Student struct {
	firstName string
	lastName string
	id int
	weight float32
	measurements map[string]int
}

// create new student: some properties initialised with value
// func newStudent(firstName, lastName string, measurements map[string]int) student {
// 	s := student{
// 		firstName: firstName,
// 		lastName: lastName,
// 		id: 0001,
// 		weight: 55.0,
// 		measurements: map[string]int{},
// 	}

// // 	return s

// // }

// func main() {
// 	s := newStudent("Oyindamola", "Abiola")

// 	fmt.Println(s)
// }

/*  
	this function has to be a reciever of Student struct-type pointer only using *Student
	to modify the instance of Student it is called on.
	Also the method parameters need to be defined, some at least.
*/
// func (s *Student) newStudent(firstName, lastName string) Student {
// 	*s = Student{
// 		// firstName: "Oyindamola",
// 		// lastName: "Abiola",
// 		firstName: firstName,
// 		lastName: lastName,
// 		id: 7,
// 		weight: 65,
// 		measurements: map[string]int{"length": 55, "burst": 38, "waist": 28, "hips": 41},
// 	}

// 	return *s
// }

// func main() {
// 	// create an instance of Student
// 	updateStudent := Student{
// 		firstName: "Doyinsola", 
// 		lastName: "Adeyinka", 
// 		id: 0, 
// 		weight: 0.0, 
// 		measurements: map[string]int{"length": 50, "burst": 38, "waist": 28},
// 	}

// 	// fmt.Println(updateStudent.newStudent())
// 	fmt.Printf("%+v\n", updateStudent.newStudent("Tolulope", "Adebayo"))
// }

// SIMPLE EXAMPLE FROM CLASS
// type Car struct {
// 	Model string
// 	Color string
// 	PlateNo int
// }

// func (c *Car) details() string {
// 	// cannot concatenate int with string
// 	return c.Model + " " + c.Color
// }

// func main() {
// 	bisola := Car{Model: "Ford", Color: "Red", PlateNo: 1505}
	
// 	fmt.Println(bisola.details())
// 	// fmt.Println(bisola.PlateNo)
// }


type DreamHome struct {
	city string
	state string
	country string
	apartment string
	floorNo int
}


// func (h DreamHome) setHome(city, country string) DreamHome {
// Initialization function, takes in what is needed as parameter
func setHome(city, state, country, apartment string, floorNo int) DreamHome { // initialize few of the struct properties
	h :=  DreamHome {
		city: city,
		state: state,
		// country: "United States",
		// apartment: "Luxury High Rise",
		country: country,
		apartment: apartment,
		floorNo: floorNo,
	}
	return h
}

// Execution function
func main() {
	// optionA := setHome("Chicago", "Illinois")
	optionB := DreamHome {
		city: "New York City", 
		state: "New York", 
		country: "United States", 
		apartment: "Luxury High Rise", 
		floorNo: 14,
	}

	fmt.Println(optionB)
}

// // Using Pointer and Deference

// type Staff struct {
// 	fullName string
// 	salary float32
// }

// func (s Staff) staffing() string {
// 	s = {

// 	}
// }