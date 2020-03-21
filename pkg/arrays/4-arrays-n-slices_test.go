package arrays

import "testing"

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
