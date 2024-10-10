package main

import (
	"fmt"
	"sort"
)

func main() {
	// Example: Sorting an integer slice
	numbers := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(numbers)
	fmt.Println("Sorted numbers:", numbers)

	// Example: Sorting a string slice
	strings := []string{"apple", "orange", "banana"}
	sort.Strings(strings)
	fmt.Println("Sorted strings:", strings)
}
