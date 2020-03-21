package arrays

//Sum receives a slice of int and returns the sum of them.
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

//SumAll sums all slice values and returns their sums inside a new slice
func SumAll(slices ...[]int) (sums []int) {
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}

	return sums
}
