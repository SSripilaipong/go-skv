package goutil

import (
	"fmt"
)

func CastOrPanic[T any](v any) T {
	result, ok := v.(T)
	if !ok {
		panic(fmt.Errorf("cast value failed"))
	}
	return result
}

func CanCast[T any](v any) bool {
	_, ok := v.(T)
	return ok
}

func May[T comparable, R any](t T, f func(T) R) (zeroResult R) {
	var zero T
	if t == zero {
		return zeroResult
	}
	return f(t)
}
