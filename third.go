package main

import (
	"fmt"
	"math"
)

// func two() {
// 	fmt.Println("Hello, World!")
// 	fmt.Println("Happy Coding!")
// 	fmt.Println("The square root of 16 is", math.Sqrt(16))
// }

// go run third.go
/*
# command-line-arguments
runtime.main_main·f: function main is undeclared in the main package
*/

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Happy Coding!")
	fmt.Println("The square root of 16 is", math.Sqrt(16))


	// constant 
	const user string = "AJ"
	
	// in go every variable must be used and shoould have a type 
	var name string

	// Just Print to show a menaing ful msg to user not LN for new line
	fmt.Print("Enter your name: ")
	// user input 
	fmt.Scan(&name)
	fmt.Println("Hello", name, "You are talking to", user)
	// & ==> pointer
	// ERORR => cannot ASSIGN to value to constant 
	//user = name

}

/*
The fmt.Scan() function is a great function for easily fetching & using user input through the command line.

But this function also has an important limitation: You can't (easily) fetch multi-word input values. Fetching text that consists of more than a single word is tricky with this function
*/
