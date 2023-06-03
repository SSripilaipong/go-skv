package clientcli

type Dependency struct {
	ConnectToServer func(string) error
}
