package main

import (
	"fmt"
	"sort"
)

func main() {
	// sorting strings

	fruits := []string{"apples", "peaches", "oranges"}

	fmt.Println(fruits)
	sort.Strings(fruits)
	fmt.Println(fruits)

	// sorting integer
	myage := []int{11, 51, 31}

	fmt.Println(myage)
	sort.Ints(myage)
	fmt.Println(myage)

	// if number are sorted
	myageSorted := []int{12, 61, 31}
	fmt.Println(myageSorted)
	sort.IntsAreSorted(myageSorted)
	fmt.Println(myageSorted)

}
