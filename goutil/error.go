package goutil

import "fmt"

func PanicUnhandledError(err error) {
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}
}
