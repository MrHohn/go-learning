package main

import "fmt"

func main() {
	for i := 1; i <= 100; i = i + 1 {
		printed := false
		if i % 3 == 0 {
			fmt.Printf("Fizz")
			printed = true
		}
		if i % 5 == 0 {
			fmt.Printf("Buzz")
			printed = true
		}
		if printed == false {
			fmt.Printf("%v", i)
		}
		fmt.Println()
	}
}