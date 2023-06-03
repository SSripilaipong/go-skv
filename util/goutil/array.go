package goutil

import "fmt"

func ElementAt[T any](array []T, index int) (T, error) {
	var zero T
	if len(array)-1 < index {
		return zero, fmt.Errorf("array index out of range")
	}
	return array[index], nil
}
