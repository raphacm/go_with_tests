package main

import (
	"fmt"
)

//Hello return 'Hello World' text
func Hello(name string) string {
	if name == "" {
		return "Hello World"
	}
	return "Hello " + name
}

func main() {
	fmt.Println(Hello("Raphael"))
}
