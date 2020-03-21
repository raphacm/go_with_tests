package arrays

//Sum receives a slice of int and returns the sum of them.
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	
	return sum
}