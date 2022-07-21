package mobile

import (
	"encoding/json"
	"errors"

	"github.com/sonr-io/sonr/bind"
	mtr "github.com/sonr-io/sonr/internal/motor"
	apiv1 "go.buf.build/grpc/go/sonr-io/motor/api/v1"
	_ "golang.org/x/mobile/bind"
)

func Init(buf []byte) ([]byte, error) {
	// Unmarshal the request
	var req apiv1.InitializeRequest
	if err := json.Unmarshal(buf, &req); err != nil {
		return nil, err
	}

	// Check if public key provided
	if req.DeviceKeyprintPub == nil {
		// Create Motor instance
		var err error
		bind.Instance, err = mtr.NewMobileMotor(req.DeviceId)
		if err != nil {
			return nil, err
		}

		// Return Initialization Response
		resp := apiv1.InitializeResponse{
			Success: true,
		}
		return json.Marshal(resp)
	}
	return nil, errors.New("Loading existing account not implemented")
}

/*
func Init(buf []byte) ([]byte, error) {
	return bind.Init(buf)
}
*/

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
