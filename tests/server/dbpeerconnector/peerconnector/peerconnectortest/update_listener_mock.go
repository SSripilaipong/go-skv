package peerconnectortest

import "go-skv/server/dbpeerconnector/peerconnectorcontract"

type UpdateListenerMock struct{}

func (u *UpdateListenerMock) OnDataUpdate(update peerconnectorcontract.DataUpdate) {
	//TODO implement me
	panic("implement me")
}

var _ peerconnectorcontract.UpdateListener = &UpdateListenerMock{}
