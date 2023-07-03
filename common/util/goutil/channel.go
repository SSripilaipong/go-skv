package goutil

import (
	"fmt"
	"time"
)

func ReceiveNoBlock[T any](dataChan <-chan T) (T, bool) {
	var zero T
	select {
	case data := <-dataChan:
		return data, true
	default:
		return zero, false
	}
}

func ReceiveWithTimeout[T any](dataChan <-chan T, timeout time.Duration) (T, bool) {
	var zero T
	select {
	case data := <-dataChan:
		return data, true
	case <-time.After(timeout):
		return zero, false
	}
}

func ReceiveWithTimeoutOrPanic[T any](dataChan chan T, timeout time.Duration) T {
	message, ok := ReceiveWithTimeout(dataChan, timeout)
	if !ok {
		panic(fmt.Errorf("unexpected error"))
	}
	return message
}

func SendWithTimeout[T any](dataChan chan T, data T, timeout time.Duration) bool {
	select {
	case dataChan <- data:
		return true
	case <-time.After(timeout):
		return false
	}
}

func SendWithTimeoutOrPanic[T any](dataChan chan T, data T, timeout time.Duration) {
	if !SendWithTimeout(dataChan, data, timeout) {
		panic(fmt.Errorf("unexpected error"))
	}
}
