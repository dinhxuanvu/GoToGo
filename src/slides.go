package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(s[1])
	fmt.Println(len(s))

	s = append(s, "d")
	fmt.Println(s)
	s = append(s, "e")
	fmt.Println(s)

	l := s[1:4]
	fmt.Println(l)
	h := s[:5]
	fmt.Println(h)
	k := s[2:]
	fmt.Println(k)

	t := []string{"g", "h", "i"}
	fmt.Println(t)

	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		twoD[i] = make([]int, i+1)
		for j := 0; j < (i + 1); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
