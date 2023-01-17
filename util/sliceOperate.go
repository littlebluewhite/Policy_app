package util

import "sort"

func SliceCopy[T any](dir []T) []T {
	result := make([]T, 0, len(dir))
	for _, v := range dir {
		result = append(result, v)
	}
	return result
}

func SliceRemove[T any](dir []T, index []int) []T {
	sort.Slice(index, func(i, j int) bool {
		return index[i] > index[j]
	})
	copyDir := SliceCopy[T](dir)
	for _, v := range index {
		copyDir = append(copyDir[:v], copyDir[v+1:]...)
	}
	return copyDir
}
