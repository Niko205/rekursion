package search

import "fmt"

func ExampleFindSorted() {
	l1 := []int{1, 3, 4, 6, 10, 21, 38}

	fmt.Println(FindSorted(l1, 10))
	fmt.Println(FindSorted(l1, 3))
	fmt.Println(FindSorted(l1, 5))

	// Output:
	// 4
	// 1
	// -1

}
