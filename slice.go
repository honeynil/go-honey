package honey

import (
	"cmp"
)

// Merge объединяет несколько слайсов в один.
func MergeSlice[T any](slices ...[]T) []T {
	total := 0
	for _, s := range slices {
		total += len(s)
	}
	result := make([]T, 0, total)
	for _, s := range slices {
		result = append(result, s...)
	}

	return result

}

// MergeSliceInPlace объединяет несколько  слайсов в один in place.
func MergeSliceInPlace[T any](a *[]T, others ...[]T) {
	for _, s := range others {
		*a = append(*a, s...)
	}
}

// MergeSortedSlice объединяет 2 sorted слайса в один.
func MergeSortedSlice[T cmp.Ordered](a, b []T) []T {
	result := make([]T, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}
