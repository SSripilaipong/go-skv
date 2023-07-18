package recordupdater

import (
	"go-skv/common/util/goutil"
	"go-skv/tests"
)

func sendWithTimeout(actor chan<- any, message any) bool {
	return goutil.SendWithTimeout(actor, message, defaultTimeout)
}

func waitForMessageWithTimeout[T any](actor <-chan any, t T) (T, bool) {
	return tests.WaitForMessageWithTimeout(actor, t, defaultTimeout)
}
