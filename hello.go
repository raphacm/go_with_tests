package main

import (
	"fmt"
)

const spanish = "Spanish"

//Hello return 'Hello World' text
func Hello(name, language string) string {
	var greetPrefix = "Hello "

	if language == spanish {
		greetPrefix = "Hola "
	}

	if name == "" {
		name = "World"
	}

	return greetPrefix + name
}

func main() {
	fmt.Println(Hello("Raphael", "Spanish"))
}
