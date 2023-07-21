package recordreplicator

import (
	"go-skv/common/util/goutil"
	"go-skv/tests"
)

func sendWithTimeout(actor chan<- any, message any) bool {
	defer close(actor)
	return goutil.SendWithTimeout(actor, message, defaultTimeout)
}

func waitForMessageWithTimeout[T any](actor <-chan any) (T, bool) {
	return tests.WaitForMessageWithTimeout[T](actor, defaultTimeout*10)
}
