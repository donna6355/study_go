package main //code executed as an application must be in a main package

import (
	"fmt"
	"isaac" //1. import package
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	msg, err := isaac.Hello("Mirrr")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(msg)
}

//1. import package
//2. go mod edit -replace isaac=../practice
//   replace local dir for the package as it is not published
//3. go mod tidy
//   add new module requirements and synchronize version
