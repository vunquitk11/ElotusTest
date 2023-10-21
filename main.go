package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 2, 1}
	b := []int{3, 2, 1, 4, 7}
	c := findLength(a, b)
	fmt.Println("RESULT", c)
}
