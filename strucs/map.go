package main

import "fmt"

var m map[string]Student

func main() {
	m = make(map[string]Student)
	n := m
	m["looloo"] = Student{
		"pourya", "Jafarzadeh", 33,
	}

	fmt.Println(m["looloo"], len(n))
}
