package main

import "fmt"

const Pi = 3.14 //Constants cannot be declared using the := syntax.

var c, python, java bool        //declare without initializer
var i, j int = 1, 2             // declare with initializer
var d, javaScript = true, "no!" // with initializer, can skip type

// varName first, then type last!!
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	var empty string //"" empty string
	var k int        //0, false, default value
	fmt.Println(empty, k, c, python, java)
	fmt.Println("Hello Isaac:)")
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println(i, j, d, javaScript)

	hello := "hello" // short var declaration only inside function
	fmt.Println(hello)

	fmt.Println("Happy", Pi, "Day")
}

/*
Basic Types
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/
