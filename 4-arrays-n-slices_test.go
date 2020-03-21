package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertEquals := func(t *testing.T, expected, received int) {
		t.Helper()

		if expected != received {
			t.Errorf("Expected %d but got %d", expected, received)
		}
	}

	t.Run("should return 16", func(t *testing.T) {
		sum := Sum([]int{8, 8})
		expected := 16

		assertEquals(t, expected, sum)
	})
}

func TestSumAll(t *testing.T) {
	assertEquals := func(t *testing.T, expected, received []int) {
		t.Helper()

		if !reflect.DeepEqual(expected, received) {
			t.Errorf("Expected %v but got %v", expected, received)
		}
	}

	t.Run("should return an slice of sums other slices, in this case [20]", func(t *testing.T) {
		sum := SumAll([]int{2, 2, 2, 2, 2}, []int{5, 4, 1})
		expected := []int{10, 10}

		assertEquals(t, expected, sum)
	})
}
