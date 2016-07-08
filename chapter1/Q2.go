package main

import "fmt"

func main() {
	// for i := 0; i < 10; i += 1 {
	// 	fmt.Printf("%d\n", i)
	// }

// 	i := 0
// again:
// 	fmt.Printf("%d\n", i)
// 	if i < 9 {
// 		i += 1
// 		goto again
// 	}

	var array1 [10]int
	for i := 0; i < 10; i += 1 {
		array1[i] = i
	}

	fmt.Printf("%v\n", array1)
}