package dbmanagertest

import "context"

type DbStorageMock struct {
	Start_ctx     context.Context
	Stop_IsCalled bool
}

func (s *DbStorageMock) Start(ctx context.Context) error {
	s.Start_ctx = ctx
	return nil
}

func (s *DbStorageMock) Stop() error {
	s.Stop_IsCalled = true
	return nil
}
