package main

import "fmt"

func main() {
	st1 := Student{name: "Pourya", lastname: "Jafarzadeh", age: 33}
	// Assign by value
	st2 := st1
	st2.name = "Nima"
	fmt.Println(st1.name, st2.name)

}
