package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("basic way of adding in an array", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
func TestSum2(t *testing.T) {
	t.Run("refactored way of summing an array", func(t *testing.T) {
		numbers := []int{1, 1, 1, 2, 2}

		got := Sum2(numbers)
		want := 7

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("Can sum any size of array", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum2(numbers)
		want := 6

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})

}
