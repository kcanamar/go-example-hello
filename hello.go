// declare a main package
// * Go code executed as an application must be in a main package

package main

import (
	// import fmt for handling inout and output text for printing to the console
	"fmt"

	// import the greetings package to gain access to the Hello func
	"example.com/greetings"

)

func main() {
	// create a greeting by dot chaining the Hello method from the greetings package
	// * := shorthand vs. var message string message = greetings.Hello()
	message := greetings.Hello("Kyle")
	// print the message to the console
	fmt.Println(message)
}