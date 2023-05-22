package grpcTest

import "fmt"

func LocalAddress(port int) string {
	return fmt.Sprintf("localhost:%d", port)
}
