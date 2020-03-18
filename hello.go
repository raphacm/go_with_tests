package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"

//Hello return 'Hello World' text
func Hello(name, language string) string {
	var greetPrefix = "Hello "

	if language == spanish {
		greetPrefix = "Hola "
	}

	if language == french {
		greetPrefix = "Bonjour "
	}

	if name == "" {
		name = "World"
	}

	return greetPrefix + name
}

func main() {
	fmt.Println(Hello("Raphael", "Spanish"))
}
