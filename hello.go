// declare a main package
// * Go code executed as an application must be in a main package

package main

import (
	// import fmt for handling inout and output text for printing to the console
	"fmt"
	// import log package to output error information
	"log"

	// import the greetings package to gain access to the Hello func
	"example.com/greetings"

)

func main() {
	/*
		! Logger Setup
		Set the properties of the predefined Logger
		Include the log entry prefix and a flag to disable printing, the time, source file, and line number.
	*/
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// create a greeting by dot chaining the Hello method from the greetings package
	// * := shorthand vs. var message string message = greetings.Hello()
	// message := greetings.Hello("Kyle")

	//? Request a greeting message
	message, err := greetings.Hello("Kyle")

	// If an error was returned, print it to the console and exit program with log packages Fatal function
	if err != nil {
		log.Fatal(err)
	}

	//! If no error was returned 
	// print the message to the console
	fmt.Println(message)
}