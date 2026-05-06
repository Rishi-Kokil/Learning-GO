package main

import "fmt"

func LoopsDemo() {
	fmt.Println("Loops Demo")

	// Classic For Loop 
	
	// Structure of this for loop
	// for initialization, condition, termination {}
	for i:= 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While style For loop
	a := 0
	for a < 5 {
		fmt.Println(a)
		a++;
	}

	// Infinite Loop
	j := 0
	for {
		fmt.Println("hello")
		if(j > 10) {
			break 	// this is how you can terminate the loop (finite or infinite) using break statement
		}
		j++
	}


	// Continue operation to skip the step
	k := 0
	for k < 5 {
		if k == 3 {
			continue
		}
		fmt.Println(k)
		k++;
	}

	// Labeled loops (very important, unique to Go)
	// This is a nested loop, and we want to break out of the outer loop when a certain condition is met in the inner loop

	outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer // This will break out of the outer loop when i is 1 and j is 1
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}

	// Iterating over a Map

	tempMap := map[string]int {
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5
	}

	for key , value := range tempmap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	} 

	// Alternate way to access the value is map[key], similar to the way we access Array elements
	for key := range tempMap {
		fmt.Printf("Key: %s, Value: %d\n", key, tempMap[key])
	}

	

}