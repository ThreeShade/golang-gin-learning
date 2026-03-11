package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = append(s, 4, 5, 6, 1, 2, 3, 7)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = s[0:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
