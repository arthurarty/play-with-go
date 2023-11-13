package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
}

// badNaming is an exported variable with a name that does not follow the camelCase convention.
// var badNaming int

// BadFunction is a function with a name that does not follow the CamelCase convention.
// func BadFunction() {
//     fmt.Println("This function has a bad name.")
// }
