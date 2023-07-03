package goutil

import "fmt"

func PanicUnhandledError(err error) {
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}
}

func WillPanicUnhandledError(f func() error) func() {
	return func() {
		err := f()
		if err != nil {
			panic(fmt.Errorf("unhandled error: %f", err))
		}
	}
}
