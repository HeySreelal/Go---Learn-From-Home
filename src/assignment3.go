package main

import "fmt"
func main(){
	n1 := 0
	n2 := 0
	op := "+"

	fmt.Printf("Enter first number: ")
	fmt.Scan(&n1)
	fmt.Printf("Enter second number: ")
	fmt.Scan(&n2)
	fmt.Printf("Enter an operator: (+ - * /): ")
	fmt.Scanf("%s", &op)

	fmt.Printf("Result: ")
	switch op {
	case "+": fmt.Printf("%v\n", add(n1, n2))
	case "-": fmt.Printf("%v\n", sub(n1, n2))
	case "*": fmt.Printf("%v\n", mult(n1, n2))
	case "/": fmt.Printf("%v\n", div(float64(n1), float64(n2)))
	}
}

// Funcs 

func add(n1 int, n2 int) int {
	return n1 + n2
}

func sub(n1 int, n2 int) int {
	return n1 - n2
}

func mult(n1 int, n2 int) int {
	return n1 * n2
}

func div(n1 float64,  n2 float64) float64 {
	return n1 / n2
}
