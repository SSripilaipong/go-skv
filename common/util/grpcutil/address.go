package grpcutil

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func LocalAddress(port int) string {
	return fmt.Sprintf("localhost:%d", port)
}

func GetPortFromAddress(addr net.Addr) (int, error) {
	tokens := strings.Split(addr.String(), ":")
	return strconv.Atoi(tokens[len(tokens)-1])
}
