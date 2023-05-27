package dbstorageTest

type RecordMock struct {
	SetValue_value string
}

func (r *RecordMock) SetValue(value string) error {
	r.SetValue_value = value
	return nil
}
