package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	fmt.Println(nums)

	for _, item := range nums {
		sum += item
	}
	fmt.Println(sum)

	str := make([]string, 3)
	str[0] = "zero"
	str[1] = "one"
	str[2] = "two"

	for i, item := range str {
		fmt.Println(i, item)
	}

	kv := map[string]int{"one": 1, "two": 2, "three": 3}

	for key, val := range kv {
		fmt.Println(key, val)
	}
}
