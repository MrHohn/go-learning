// // 1
// package main

// import "fmt"

// func main() {
// 	repeat := 1
// 	for repeat <= 10 {
// 		for i := 0; i < repeat; i++ {
// 			fmt.Printf("A")
// 		}
// 		fmt.Println()
// 		repeat++
// 	}
// }


// // 2
// package main

// import "fmt"
// import "unicode/utf8"

// func main() {
// 	input := "asSASA ddd dsjkdsjs dk"
// 	fmt.Printf("String: %s\nLength: %d, Runes: %d\n", input, len([]byte(input)), utf8.RuneCount([]byte(input)))
// }


// // 3
// package main

// import "fmt"

// func main() {
// 	input := "asSASA ddd dsjkdsjs dk"
// 	r := []rune(input)
// 	copy(r[4:], []rune("abc"))
// 	// copy(r[4:4+3], []rune("abc"))
// 	fmt.Printf("Before: %s\n", input)
// 	fmt.Printf("After: %s\n", string(r))
// }

// 4
package main

import "fmt"

func main() {
	input := "foobar"
	input_rune := []rune(input)
	for i, j := 0, len(input) - 1; i < j; i, j = i + 1, j - 1 {
		input_rune[i], input_rune[j] = input_rune[j], input_rune[i]
	}
	fmt.Printf("%s\n", string(input_rune))
}
