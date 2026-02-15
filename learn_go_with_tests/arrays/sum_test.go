package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum a slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 21

		if got != want {
			t.Errorf("Expected '%d', got '%d' given %v", want, got, numbers)
		}
	})
}
