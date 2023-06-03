package clientcli

type dependency struct {
	ConnectToServer func(string) error
}
