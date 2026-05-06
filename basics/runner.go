package main

import "fmt"

func main() {
	VariableDemo()
	LoopsDemo()


	intSlice := []int {12,3,45,6,7,87,3,22,46,66}

	count := 0

	for _,value := range  intSlice {
		if(value % 2 == 0) { 
			count++
		}
	}

	fmt.Println("Even count:", count)

}
