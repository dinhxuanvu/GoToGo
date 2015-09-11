package main

import "fmt"

func sum(nums ...int) (int, int) {
	count := 0
	total := 0
	for _, item := range nums {
		count++
		total += item
	}
	return count, total
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum(nums...))
}
