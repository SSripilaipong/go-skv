package recordreplicator

import (
	"go-skv/common/test"
	"go-skv/common/util/goutil"
)

func sendWithTimeout(actor chan<- any, message any) bool {
	defer close(actor)
	return goutil.SendWithTimeout(actor, message, defaultTimeout)
}

func waitForMessageWithTimeout[T any](actor <-chan any) (T, bool) {
	return test.WaitForMessageWithTimeout[T](actor, defaultTimeout*10)
}
