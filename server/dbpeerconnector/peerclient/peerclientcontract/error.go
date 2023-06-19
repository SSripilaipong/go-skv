package peerclientcontract

type ConnectionError struct{}

func (e ConnectionError) Error() string {
	return "connection error"
}

func (e ConnectionError) Is(d error) bool {
	switch d.(type) {
	case ConnectionError:
		return true
	}
	return false
}
