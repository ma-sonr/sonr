//go:build wasm
// +build wasm

package bind

import (
	"encoding/json"
	"errors"

	mtr "github.com/sonr-io/sonr/internal/motor"
	apiv1 "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

var (
	errWalletExists    = errors.New("mpc wallet already exists")
	errWalletNotExists = errors.New("mpc wallet does not exist")
)

var instance *mtr.MotorNode

func Init(buf []byte) ([]byte, error) {
	// Unmarshal the request
	var req apiv1.InitializeRequest
	if err := json.Unmarshal(buf, &req); err != nil {
		return nil, err
	}

	// Check if public key provided
	if req.DeviceKeyprintPub == nil {
		// Create Motor instance
		instance = mtr.EmptyMotor(req.DeviceId)

		// Return Initialization Response
		resp := apiv1.InitializeResponse{
			Success: true,
		}
		return json.Marshal(resp)
	}
	return nil, errors.New("Loading existing account not implemented")
}

func CreateAccount(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}
	if res, err := instance.CreateAccount(buf); err == nil {
		return json.Marshal(res)
	} else {
		return nil, err
	}
}

func Login(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	if res, err := instance.Login(buf); err == nil {
		return json.Marshal(res)
	} else {
		return nil, err
	}
}

// Address returns the address of the wallet.
func Address() string {
	if instance == nil {
		return ""
	}
	addr, err := instance.Wallet.Address()
	if err != nil {
		return ""
	}
	return addr
}

// Balance returns the balance of the wallet.
func Balance() int {
	return int(instance.Balance())
}

// DidDoc returns the DID document as JSON
func DidDoc() ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}
	buf, err := instance.DIDDocument.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return buf, nil
}
