package main

import "fmt"

type  Bill struct {
	customerName string
	items map[string]float32
	tip float32
}

func newBill(name string) Bill {
	b := Bill {
		customerName: name,
		items: map[string]float32{"cake": 1.55, "pie": 0.77},
		tip: 0,
	}

	return b
}

func (b Bill) updateBill() string {
	fs := "Bill Breakdown: \n"
	var total float32 = 0

	// list item
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) 
		total += v
	}

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)
	return fs
}

func main() {
	b := newBill("Mauri")

	fmt.Println(b.updateBill())
}