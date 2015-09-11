package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func multiplus(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Println(plus(1, 1))
	fmt.Println(multiplus(1, 2, 3))
	fmt.Println(minus(2, 1))
	fmt.Println(multiply(2, 2))
	fmt.Println(divide(1, 0))
	fmt.Println(divide(8, 2))
}
