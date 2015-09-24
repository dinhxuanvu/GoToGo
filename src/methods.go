package main

import "fmt"

func main() {
	r := rect{10, 20}
	rp := &r
	fmt.Println(r.area())
	fmt.Println(r.perim())
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perim() int {
	return 2 * (r.height + r.width)
}
