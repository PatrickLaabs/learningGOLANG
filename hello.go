package main

import (
	"fmt"
)

// Hello exported for testing
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
