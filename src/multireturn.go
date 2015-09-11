package main

import "fmt"

func sum(a, b, c, d int) (int, int) {
	return 4, (a + b + c + d)
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
}
