package main

/*
Question:
Write a program that:

1. Declares a slice of integers
2. Uses a loop to count how many numbers are even
3. Print:
   Even count: X
4. Use:
   - range
   - if
   - % operator
*/

import "fmt"

func main() {
	intSlice := []int{12, 3, 45, 6, 7, 87, 3, 22, 46, 66}

	count := 0

	for _, value := range intSlice {
		if value%2 == 0 {
			count++
		}
	}

	fmt.Println("Even count:", count)
}