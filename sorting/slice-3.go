package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	// Sort by age using sort.Slice
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("Sorted by age:", people)
}
