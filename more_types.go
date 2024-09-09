package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func pointer() {
	i, j := 42, 2701
	p := &i         // pointer address to i
	fmt.Println(*p) // cannot print pointer address, need to dereference it!
	*p = 21         // allocate new val thru pointer address
	fmt.Println(i)  // val has been changed

	p = &j         // override pointer to j's address
	*p = *p / 37   // allocate new val thru pointer address
	fmt.Println(j) // val has been changed
}

type Cat struct { //struct is a collection of fields
	name string
	age  int
}

func meow() {
	fmt.Println(Cat{"Isaac", 4})
	cat := Cat{"moomin", 7}
	cat.age = 8 //Struct fields are accessed using a dot
	fmt.Println(cat.name)

	p := &cat
	//(*p).age = 9 // that notation is cumbersome
	p.age = 9 // struct pointer is allowed without the explicit dereference
	p.name = "cute moomin"

	cat2 := Cat{name: "Kkomi"} //age:0 is implicit
	cat3 := Cat{}              //name:"", age:0
	fmt.Println(cat2.name)
	fmt.Println(cat2.age)
	fmt.Println(cat3)
}

func array() { // arrays cannot be resized
	var cats [3]Cat //an array of n values of type T
	cats[0] = Cat{name: "Isaac"}
	cats[1] = Cat{}
	cats[2] = Cat{name: "Moomin"}
	fmt.Println(cats)

	bigCats := [10]Cat{Cat{}, Cat{}, Cat{}, Cat{}, Cat{}, Cat{}, Cat{}, Cat{}, Cat{}, Cat{}}
	fmt.Println(bigCats)
}

func slice() { // a dynamically-sized
	primes := []int{2, 3, 5, 7, 11, 13}

	var s []int = primes[0:3] //a[low : high]
	fmt.Println(s)            //2,3,4 includes the first element, but excludes the last one

	//Slices are like references to arrays
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] // names[:2]  low default 0
	b := names[1:4] // names[1:]  hight default the length
	fmt.Println(a, b)

	//Changing the elements of a slice modifies the corresponding elements of its underlying array
	b[0] = "XXX"
	fmt.Println(a, b)  //[John XXX] [XXX George]
	fmt.Println(names) //[John XXX George Ringo]

	//The length of a slice is the number of elements it contains.
	//The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

	var sl []int
	fmt.Println(sl, len(sl), cap(sl))

	//create a slice with make
	x := make([]int, 5)
	printSlice("x", x)

	y := make([]int, 0, 5)
	printSlice("y", y)

	//append to a slice
	a = append(names[:], "Isaac")
	fmt.Println(a)
}

func iterateRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow { //idx, value
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//skip value
	// for i, _ := range pow //for i := range pow
	//skip index
	// for _, value := range pow
}
