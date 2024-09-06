package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func Sqrt(x float64) float64 {
	z := float64(1)
	fmt.Printf("Given num is %v\n", x)
	for i := 0; i < 15; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("#%v try : %v\n", i+1, z)
	}
	return z

}
func flow() {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	for sum < 10 { //The init and post statements are optional. => same for the while in other lang
		sum += sum
	}

	fmt.Println(sum)

	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func osCheck() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	//the break statement that is needed at the end of each case in those languages is provided automatically in Go
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func weekend() {
	fmt.Println("Where is Weekend???")
	today := time.Now().AddDate(0, 0, 1).Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("TMRRW")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func noConditionSwitch() {
	//Switch without a condition is the same as switch true.
	fmt.Println("Switch without a condition")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Still in the morning")
	case t.Hour() > 12:
		fmt.Println("Oops, in the afternoon")
	default:
		fmt.Println("Exactly at NOON!!!")
	}

}

func tryDefer() {
	defer fmt.Println("Isaac") //defers the execution of a function until the surrounding function returns.
	fmt.Println("Good")
	//Good
	//Isaac

	//Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed
	//in last-in-first-out order.
	fmt.Println("Start")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Finish")
	//Start, Finish, 9,8,7,6,5,4,3,2,1,0
}
