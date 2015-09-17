package main

import "fmt"

func main() {
	fmt.Println(fibonacci(10))
	fmt.Println(fib(10))
}

func fibonacci(num int) int {
	if num < 2 {
		return 1
	}
	var previous, current, sum int = 1, 1, 0
	for i := 2; i <= num; i++ {
		sum = previous + current
		previous = current
		current = sum
	}
	return current
}

func fib(num int) int {
	if num < 2 {
		return 1
	} else {
		return (fib(num-2) + fib(num-1))
	}
}
