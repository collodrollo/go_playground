package arrays

import (
	"slices"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, want, got []int) {
		t.Helper()
		if !slices.Equal(want, got) {
			t.Errorf("expected %v, got %v", want, got)
		}
	}

	t.Run("sum tails of a few slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, want, got)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{})
		want := []int{0, 0}

		checkSums(t, want, got)
	})
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for b.Loop() {
		SumAllTails([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	}
}

// go test -bench=.
