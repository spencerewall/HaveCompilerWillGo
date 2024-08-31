package gobyexample

// basic library for i/o and string formatting https://pkg.go.dev/fmt
import "fmt"

func main() {
	var string1 = "you can declare variables with a simple ="
	string2 := "or you can declare a variable and instatiate it with :="
	fmt.Println(string1)
	fmt.Println(string2)

	fmt.Println()

	whileCounter := 0
	for whileCounter < 5 {
		whileCounter++
		whileMsg := "go has no while loops, it's all f"
		oCounter := 0
		for oCounter < whileCounter {
			whileMsg += "o"
			oCounter++
		}
		whileMsg += "r"
		fmt.Println(whileMsg)
	}
	// but the following will error
	// string2 := "but this will error"
	fmt.Println()
	for j := 0; j < 3; j++ {
		fmt.Println("it's also got normal style for loops, but there are no parens")
	}

	fmt.Println()

	for rangeCounter := range 6 {
		fmt.Print(rangeCounter)
		fmt.Println("it's also got syntactic sugar for ranges, just like ruby!")
	}

	fmt.Println()

	fmt.Println("For conditional control, go has if/else, and case/switch statements. BUT no ternerary operators!!")

	if num := 9; num < 0 {
		fmt.Println(num, "if else statements function basically as normal, except you can have multiple statements in a single line, which is pretty cool")
	} else if num < 10 {
		fmt.Println(num, "if else statements function basically as normal, except you can have multiple statements in a single line, which is pretty cool.")
		fmt.Println("variables declared in prior conditional statements are available in later conditionals, which I can't imagine a greate use for yet")
	} else {
		fmt.Println("not here tho")
	}

	fmt.Println()
	fmt.Println("Switch statements are pretty powerful in go")
}
