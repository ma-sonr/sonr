package mobile

import (
	"github.com/sonr-io/sonr/bind"
	_ "golang.org/x/mobile/bind"
)

func Init(buf []byte) ([]byte, error) {
	return bind.Init(buf)
}

func CreateAccount(buf []byte) ([]byte, error) {
	return bind.CreateAccount(buf)
}

func Login(buf []byte) ([]byte, error) {
	return bind.Login(buf)
}

func Address() string {
	return bind.Address()
}

func Balance() int {
	return bind.Balance()
}

func DidDoc() ([]byte, error) {
	return bind.DidDoc()
}
