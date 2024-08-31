package test

import (
	"fmt"
	"log"
	"reflect"
)

func Assert(condition bool, message string) {
	if !condition {
		log.Fatalf("Assertion failed: %s", message)
	}
}

func AssertEqual(any1 any, any2 any) {
	if !reflect.DeepEqual(any1, any2) {
		log.Fatalf("Assertion failed: %s", fmt.Sprintf("%s does not equal %s", any1, any2))
	}
}

func main() {
	value := 5
	Assert(value == 10, "value should be 10")
	fmt.Println("This will not print if the assertion fails")
}
