package client

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func parseGrpcError(err error) (error, error) {
	if e, ok := status.FromError(err); ok {
		if e.Code() == codes.Canceled {
			return CancelledError{}, nil
		}
	}
	return nil, err
}

type CancelledError struct{}

func (CancelledError) Error() string {
	return "context cancelled"
}
