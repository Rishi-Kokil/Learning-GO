package main

import "fmt"

func CommandLineArgumentsDemo() {
	fmt.Println("Command Line Arguments Demo")
	// os.Args is a slice of strings that contains the command line arguments passed to the program
	// The first element (os.Args[0]) is the name of the program itself
	// The subsequent elements are the actual arguments passed by the user
	fmt.Println("Command Line Arguments: ", os.Args)

	// Using loop to go through all the command line arguments
	// using the Range keyword to iterate over the slice of command line arguments
	// _ is the index, and since we don't need the index, we can use _ to ignore it
	for _, value := range os.Args {
		fmt.Println(value)
	}

	// Generally we skip the first argument (the program name) and start from the second argument
	
	for _, value := range os.Args[1:] {
		fmt.Println(value)
	}
}