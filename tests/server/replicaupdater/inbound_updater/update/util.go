package update

import "go-skv/common/util/goutil"

func sendWithTimeout(updater chan<- any, message any) bool {
	return goutil.SendWithTimeout[any](updater, message, defaultTimeout)
}
