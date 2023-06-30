package peerconnectorcontract

type CannotConnectToPeerError struct{}

func (CannotConnectToPeerError) Error() string {
	return "cannot connect to peer"
}
