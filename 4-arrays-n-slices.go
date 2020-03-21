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
func SumAll(slices ...[]int) []int {
	lengthOfSlices := len(slices)
	sums := make([]int, lengthOfSlices)

	for i, slice := range slices {
		for _, number := range slice {
			sums[i] += number
		}
	}

	return sums
}
