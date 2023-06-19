package peerconnectorcontract

type UpdateListener interface {
	OnDataUpdate(update DataUpdate)
}

type DataUpdate struct{}
