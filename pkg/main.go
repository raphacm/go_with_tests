package main

import (
	"fmt"

	"github.com/raphacm/go_with_tests/pkg/greeting"
	"github.com/raphacm/go_with_tests/pkg/integers"
	"github.com/raphacm/go_with_tests/pkg/iteration"
)

func main() {
	fmt.Println(greeting.Hello("", ""))
	fmt.Println(greeting.Hello("Raphael", ""))
	fmt.Println(greeting.Hello("Raphael", "Spanish"))
	fmt.Println(greeting.Hello("Raphael", "French"))
	fmt.Println(greeting.Hello("Raphael", "Italian"))
	fmt.Println(integers.Add(1, 1))
	fmt.Println(iteration.Repeat("z", 100))
}
