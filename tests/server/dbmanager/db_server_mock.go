package dbmanagerTest

type DbServerMock struct {
	Start_IsCalled bool
}

func (s *DbServerMock) Start() error {
	s.Start_IsCalled = true
	return nil
}
