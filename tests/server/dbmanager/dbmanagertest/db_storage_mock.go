package dbmanagertest

import "context"

type DbStorageMock struct {
	Start_ctx     context.Context
	Join_IsCalled bool
}

func (s *DbStorageMock) Start(ctx context.Context) error {
	s.Start_ctx = ctx
	return nil
}

func (s *DbStorageMock) Join() error {
	s.Join_IsCalled = true
	return nil
}
