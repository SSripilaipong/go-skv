package getValue

import "context"

func contextWithDefaultTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeout)
}
