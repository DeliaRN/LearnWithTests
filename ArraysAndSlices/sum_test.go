package ArraysAndSlices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	//Multiple tests: FIRST ONE:
	t.Run("collection of 5 numbers", func(t *testing.T) {
		var numbers = []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	//Multiple tests: SECOND ONE:
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	//this won't compile as we cannot use equality with slices
	//if got != want {
	//so we use reflect.DeepEqual ("reflect" must be imported):
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	//Assigning a function to a variable, so we can take away the checking repeated code in every test
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		/*if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given", got, want)
		}*/
		checkSums(t, got, want)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		/*if !reflect.DeepEqual(got, want){
			t.Errorf("got %v eant %v", got, want)
		}*/
		checkSums(t, got, want)
	})
}
