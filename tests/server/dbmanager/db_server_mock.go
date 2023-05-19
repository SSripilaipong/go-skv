package dbmanagerTest

type DbServerMock struct {
	Start_IsCalled bool
	Stop_IsCalled  bool
}

func (s *DbServerMock) Start() error {
	s.Start_IsCalled = true
	return nil
}

func (s *DbServerMock) Stop() error {
	s.Stop_IsCalled = true
	return nil
}
