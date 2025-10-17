package honey

import (
	"reflect"
	"testing"
)

func TestMergeSlice(t *testing.T) {
	t.Run("Merge multiple slices", func(t *testing.T) {
		a := []int{1, 2}
		b := []int{3, 4}
		c := []int{5, 6}

		got := MergeSlice(a, b, c)
		want := []int{1, 2, 3, 4, 5, 6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("MergeSlice() = %v, want %v", got, want)
		}
	})

	t.Run("Merge with empty slice", func(t *testing.T) {
		a := []int{1, 2}
		b := []int{}

		got := MergeSlice(a, b)
		want := []int{1, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("MergeSlice() = %v, want %v", got, want)
		}
	})
}

func TestMergeSliceInPlace(t *testing.T) {
	t.Run("In-place merge", func(t *testing.T) {
		a := []int{1, 2}
		b := []int{3, 4}

		MergeSliceInPlace(&a, b)
		want := []int{1, 2, 3, 4}

		if !reflect.DeepEqual(a, want) {
			t.Errorf("MergeSliceInPlace() = %v, want %v", a, want)
		}
	})
}

func TestMergeSortedSlice(t *testing.T) {
	t.Run("Merge two sorted slices", func(t *testing.T) {
		a := []int{1, 3, 5, 7}
		b := []int{2, 4, 6, 8}

		got := MergeSortedSlice(a, b)
		want := []int{1, 2, 3, 4, 5, 6, 7, 8}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("MergeSortedSlice() = %v, want %v", got, want)
		}
	})

	t.Run("One slice empty", func(t *testing.T) {
		a := []int{}
		b := []int{1, 2, 3}

		got := MergeSortedSlice(a, b)
		want := []int{1, 2, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("MergeSortedSlice() = %v, want %v", got, want)
		}
	})

	t.Run("Already sorted input", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{4, 5, 6}

		got := MergeSortedSlice(a, b)
		want := []int{1, 2, 3, 4, 5, 6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("MergeSortedSlice() = %v, want %v", got, want)
		}
	})
}
