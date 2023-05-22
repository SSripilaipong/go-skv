package goutil

func ReceiveNoBlock[T any](dataChan chan T) (T, bool) {
	var zero T
	select {
	case data := <-dataChan:
		return data, true
	default:
	}
	return zero, false
}
