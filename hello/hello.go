package main // just like the way we define packages in JAVA

// This is similar to the way we import packages in JAVA. We are importing the fmt package which contains functions for formatting text, including printing to the console.
import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(quote.Go())
}
