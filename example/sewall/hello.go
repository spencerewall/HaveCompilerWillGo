package main

// basic library for i/o and string formatting https://pkg.go.dev/fmt
import "fmt"

func main() {
	var string1 = "you can declare variables with a simple ="
	string2 := "or you can declare a variable and instatiate it with :="
	fmt.Println(string1)
	fmt.Println(string2)
	
	// but the following will error
	// string2 := "but this will error"
}
