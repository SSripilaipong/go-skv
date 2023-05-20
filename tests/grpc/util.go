package grpcTest

import "fmt"

func GetAvailablePort() int {
	return 12345
}

func LocalAddress(port int) string {
	return fmt.Sprintf("localhost:%d", port)
}
