package main

import (
	"github.com/raphacm/go_with_tests/pkg/greeting"
)

func main() {
	fmt.Println(greeting.Hello("", ""))
	fmt.Println(greeting.Hello("Raphael", ""))
	fmt.Println(greeting.Hello("Raphael", "Spanish"))
	fmt.Println(greeting.Hello("Raphael", "French"))
	fmt.Println(greeting.Hello("Raphael", "Italian"))
}
