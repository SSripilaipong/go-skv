package dbmanagertest

type DbStorageMock struct {
	Start_IsCalled bool
	Stop_IsCalled  bool
}

func (s *DbStorageMock) Start() error {
	s.Start_IsCalled = true
	return nil
}

func (s *DbStorageMock) Stop() error {
	s.Stop_IsCalled = true
	return nil
}
