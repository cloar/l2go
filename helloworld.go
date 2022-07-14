package main

import "fmt"

func main() {
	fmt.Println("Hello World.")
	foo()
	fmt.Println("This is after the foo function line.")

	tester := "42"

	fmt.Println(tester)
}

func foo() {
	n, _ := fmt.Println("I'm hungry for lunch so let's eat! NOW!!!")
	fmt.Println(n)

}

// control flow
// 1. sequence
// 2. loop: iterative
// 3. conditional

// you can't have unused variables in Go.  This is CODE POLLUTION
//
