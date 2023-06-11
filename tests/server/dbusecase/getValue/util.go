package getValue

import (
	"context"
	"go-skv/util/goutil"
)

func contextWithDefaultTimeout() (context.Context, context.CancelFunc) {
	return goutil.NewContextWithTimeout(defaultTimeout)
}
