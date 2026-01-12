package main

import (
	"fmt"
	"net/http"
	"log"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// check for err
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello again, Oyindamola")

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// first check for the error while loading the form page
	if err := r.ParseForm(); err != nil {
		// fmt.Fprintf(w, "Error while loading the form", err)
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// parse in the form value
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	age := r.FormValue("age")
	career := r.FormValue("career")

	// response print out in w
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
	fmt.Fprintf(w, "Age = %s\n", age)
	fmt.Fprintf(w, "Career = %s\n", career)
}

func main() {
	// first check check for the static file
	fileServer := http.FileServer(http.Dir("./static"))

	// handle endpoints
	http.Handle("/", fileServer) // this handles the index.html
	http.HandleFunc("/form", formHandler) // this handles the form function
	http.HandleFunc("/hello", helloHandler) // this handles the hello function
	
	// print the string for starting the server
	fmt.Printf("Static Webserver is starting at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
