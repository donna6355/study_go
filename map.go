package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex //key:string, val:Vertex
// zero value of a map is nil. A nil map has no keys, nor can keys be added.

var labs = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// If the top-level type is just a type name, you can omit it from the elements of the literal.
var omit = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func makeMap() {
	m = make(map[string]Vertex)
	m["isaac"] = Vertex{1.11, 23.3}

	fmt.Println(m["isaac"])
	fmt.Println(labs)
}

func mutating() {

	var m = map[string]string{
		"Mir":    "Meow",
		"Moomin": "Shhh",
	}
	cry := m["Moomin"]
	fmt.Println("Moomin's cry is ", cry)

	//make Map can update
	// m["Mir"] = "Grrrrr"
	// fmt.Println("Mir's cry is ", m["Mir"])

	delete(m, "Mir")
	e, ok := m["Mir"]
	fmt.Println("Mir's cry is ", e, "Mir is present :", ok)
}
