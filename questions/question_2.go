package main

/*
Question:
Write a Go program that:

1. Declares a map[string]int
2. Adds 3 key-value pairs
3. Uses a range loop
4. Prints:
   key: <key>, value: <value>
5. Use := wherever possible
*/

import "fmt"

func main() {
	tempMap := map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
	}

	for key, value := range tempMap {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
}