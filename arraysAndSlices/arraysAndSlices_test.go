package arraysandslices

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}

	got := SumAll(slice1, slice2)

	want := []int{6, 6}

	checkSums(t, got, want)

}

func TestSumAllTails(t *testing.T) {
	t.Run("run for any number of slices", func(t *testing.T) {

		slice1 := []int{1, 2}
		slice2 := []int{0, 9}

		got := SumAllTails(slice1, slice2)

		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely run for empty slices", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{}

		got := SumAllTails(slice1, slice2)
		want := []int{2, 0}

		checkSums(t, got, want)
	})
}

func checkSums(t testing.TB, got []int, want []int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v got %v", want, got)
	}
}
