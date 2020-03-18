package integers

import "testing"

func TestAdder(t *testing.T) {

	assertEqual := func(t *testing.T, expected, received int) {
		t.Helper()
		if (expected != received) {
				t.Errorf("Expected %d but got %d", expected, received)
		}
	}

	t.Run("should return 4 when sum is called with 2, 2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertEqual(t, expected, sum)
	})
}