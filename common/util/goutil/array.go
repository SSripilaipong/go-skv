package goutil

import "fmt"

func ElementAt[T any](array []T, index int) (T, error) {
	var zero T
	if len(array)-1 < index {
		return zero, fmt.Errorf("array index out of range")
	}
	return array[index], nil
}

func Map[T, R any](array []T, f func(T) R) []R {
	result := make([]R, len(array))
	for i, x := range array {
		result[i] = f(x)
	}
	return result
}
