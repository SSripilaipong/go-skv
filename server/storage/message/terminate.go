package message

type Terminate struct {
	Notify chan<- struct{}
}
