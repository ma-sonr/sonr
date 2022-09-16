package common

import "log"

type MotorCallback interface {
	OnDiscover(data []byte)
	OnWalletEvent(data []byte)
}

type defaultCallback struct{}

func DefaultCallback() MotorCallback {
	return &defaultCallback{}
}

func (cb *defaultCallback) OnWalletEvent(data []byte) {
	log.Println("ERROR: MotorCallback not implemented.")
}

func (cb *defaultCallback) OnDiscover(data []byte) {
	log.Println("ERROR: MotorCallback not implemented.")
}
