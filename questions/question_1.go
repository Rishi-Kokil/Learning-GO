package main

/*
Question:
Write a small Go program that:

1. Uses a module-style import path (assume it exists)
2. Has a main function
3. Declares a slice of integers
4. Uses a range loop
5. Prints:
   - index
   - value
6. Uses short variable declaration (:=)
*/

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for index, value := range nums {
		fmt.Printf("index=%d, value=%d\n", index, value)
	}
}