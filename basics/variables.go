package main

import "fmt"

func VariableDemo() {
	// Method 1
	var a int
	fmt.Println("A: ", a) // when we use comma the data is printed after the string

	// Method 2
	var b int = 10
	fmt.Printf("B: %d", b) // %d is used to print the integer value inside a string, its like injecting it there
	// not its Printf and not Println because we are using format specifier, and Printf does not add a new line at the end of the output

	// Method 3
	var c = 20
	fmt.Println("C: ", c)

	// Method 4
	d := 30
	fmt.Println("D: ", d)

	// Multiple variable declaration
	var e, f, g = 40, "Hello", 60
	fmt.Println("E: ", e)
	fmt.Println("F: ", f)
	fmt.Println("G: ", g)

	// Block declaration
	var (
		j int     = 70
		k string  = "block"
		l float32 = 3.14
	)
	fmt.Println("J: ", j)
	fmt.Println("K: ", k)
	fmt.Println("L: ", l)

	// ALL THE BASIC DATA TYPES
	// Integer types
	var i8 int8 = 8
	var i16 int16 = 16
	var i32 int32 = 32
	var i64 int64 = 64

	// Unsigned integers
	var u uint = 100
	var u8 uint8 = 8
	var u16 uint16 = 16
	var u32 uint32 = 32
	var u64 uint64 = 64

	// Floating point
	var f32 float32 = 3.14
	var f64 float64 = 6.28

	// Boolean
	var isGoAwesome bool = true

	// String
	var str string = "Go is simple"

	// Byte (alias for uint8)
	var by byte = 'A'

	// Rune (alias for int32, used for Unicode)
	var r rune = '₹'

	fmt.Println(i8, i16, i32, i64)
	fmt.Println(u, u8, u16, u32, u64)
	fmt.Println(f32, f64)
	fmt.Println(isGoAwesome)
	fmt.Println(str)
	fmt.Println(by)
	fmt.Println(r)

	// Complex types

	//Array
	var arr [3]int = [3]int{1, 2, 3}

	//Slice
	slice := []int{4, 5, 6}

	// Map
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "John", Age: 25}

	// Pointer
	x := 100
	ptr := &x

	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(m)
	fmt.Println(p)
	fmt.Println(ptr, *ptr)

	// -------------------------------
	// 4. Constants
	// -------------------------------

	const pi = 3.14159
	const appName string = "GoApp"

	fmt.Println(pi, appName)

	// NOTE : Go has zero values (int = 0, string = "", bool = false)

}
