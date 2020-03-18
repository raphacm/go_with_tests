package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertEquals := func(t *testing.T, expected, received string) {
		t.Helper()
		if expected != received {
			t.Errorf("Expected %q but got %q", expected, received)
		}
	}

	t.Run("should repeat 5 times the letter 'a'", func(t *testing.T) {
		got := Repeat("a", 5)
		expected := "aaaaa"

		assertEquals(t, expected, got)
	})

	t.Run("should repeat letter N times when times parameter is passed", func(t *testing.T) {
		got := Repeat("b", 15)
		expected := "bbbbbbbbbbbbbbb"

		assertEquals(t, expected, got)
	})

	t.Run("should return empty string if times parameter is less than 0", func(t *testing.T) {
		got := Repeat("z", -10)
		expected := ""

		assertEquals(t, expected, got)
	})
}
