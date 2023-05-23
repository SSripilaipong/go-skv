package goutil

import "time"

func ReceiveNoBlock[T any](dataChan chan T) (T, bool) {
	var zero T
	select {
	case data := <-dataChan:
		return data, true
	default:
		return zero, false
	}
}

func ReceiveWithTimeout[T any](dataChan chan T, timeout time.Duration) (T, bool) {
	var zero T
	select {
	case data := <-dataChan:
		return data, true
	case <-time.After(timeout):
		return zero, false
	}
}
