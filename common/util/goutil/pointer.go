package goutil

func Pointer[T any](v T) *T {
	return &v
}
