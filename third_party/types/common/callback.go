package common

import "log"

type MotorCallback interface {
	OnDiscover(data []byte)
	OnMotorEvent(msg string, isDone bool)
}

type defaultCallback struct{}

func DefaultCallback() MotorCallback {
	return &defaultCallback{}
}

func (cb *defaultCallback) OnMotorEvent(msg string, isDone bool) {
	log.Printf("message: %s done: %t", msg, isDone)
}

func (cb *defaultCallback) OnDiscover(data []byte) {
	log.Println("ERROR: MotorCallback not implemented.")
}

func (cb *defaultCallback) OnWalletCreated(ok bool) {
	log.Println("ERROR: MotorCallback not implemented.")
}

func (d defaultCallback) OnMotorEvent(msg string, done bool) {
	log.Println("ERROR: MotorCallback not implemented.")
}
