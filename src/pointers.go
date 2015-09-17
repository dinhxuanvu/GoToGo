package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(ival *int) {
	*ival = 0
}

func main() {
	i := 5
	fmt.Println(i)
	zeroval(i)
	fmt.Println(i)
	zeroptr(&i)
	fmt.Println(i)
	fmt.Println(&i)
}
