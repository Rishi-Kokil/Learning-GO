package main
/*
Write a program that:

Takes a slice of integers
Creates a map[int]int
key → number
value → frequency (count how many times it appears)

Print the map like:

3 -> 2
7 -> 1
*/

func main() {
	intSlice := []int {12,3,45,6,7,87,3,22,46,66}
	frequencyMap := make(map[int]int)

	for key := range intSlice {
		frequencyMap[intSlice[key]]++
	}

	for key, value := range frequencyMap {
		println(key, "->", value)
	}

}