package main

import "fmt"

func main() {
	numbers := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Input numbers: %v\n", numbers)
	
	if len(numbers) == 0 {
		return
	}

	sum := 0.0
	for _, v := range numbers {
		sum += v
	}
	sum /= float64(len(numbers))
	fmt.Printf("The average is: %v\n", sum)
}