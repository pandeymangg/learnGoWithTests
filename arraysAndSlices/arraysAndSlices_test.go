package arraysandslices

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		got := Sum(numbers)
		want := 21

		if got != want {
			t.Errorf("want %d, got %d", want, got)
		}
	})
}
