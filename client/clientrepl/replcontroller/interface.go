package replcontroller

type Interface interface {
	Connect(address string) (err error)
	Input(s string) (string, error)
}
