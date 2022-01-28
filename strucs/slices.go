package main

import "fmt"

func main() {

	// Define the slice with make
	//a := make([]int,5)
	var a = make([]int, 5, 10)

	// Using append to add in slices
	c := []int{}
	c = append(c, -4)

	a[0] = 3
	a = append(a, -12)

	// slices are pointers
	b := a
	b[1] = 9
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b[0:])
	fmt.Println(c)
}
