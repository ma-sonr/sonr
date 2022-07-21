package bind

import (
	"encoding/json"
	"errors"

	mtr "github.com/sonr-io/sonr/internal/motor"
)

var (
	errWalletExists    = errors.New("mpc wallet already exists")
	errWalletNotExists = errors.New("mpc wallet does not exist")
)

var Instance *mtr.MotorNode

func CreateAccount(buf []byte) ([]byte, error) {
	if Instance == nil {
		return nil, errWalletNotExists
	}
	if res, err := Instance.CreateAccount(buf); err == nil {
		return json.Marshal(res)
	} else {
		return nil, err
	}
}

func Login(buf []byte) ([]byte, error) {
	if Instance == nil {
		return nil, errWalletNotExists
	}

	if res, err := Instance.Login(buf); err == nil {
		return json.Marshal(res)
	} else {
		return nil, err
	}
}

// Address returns the address of the wallet.
func Address() string {
	if Instance == nil {
		return ""
	}
	addr, err := Instance.Wallet.Address()
	if err != nil {
		return ""
	}
	return addr
}

// Balance returns the balance of the wallet.
func Balance() int {
	return int(Instance.Balance())
}

// DidDoc returns the DID document as JSON
func DidDoc() ([]byte, error) {
	if Instance == nil {
		return nil, errWalletNotExists
	}
	buf, err := Instance.DIDDocument.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return buf, nil
}
