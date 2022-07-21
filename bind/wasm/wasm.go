//go:build wasm
// +build wasm

package main

import (
	"encoding/json"
	"errors"
	"syscall/js"

	"github.com/sonr-io/sonr/bind"
	mtr "github.com/sonr-io/sonr/internal/motor"
	apiv1 "go.buf.build/grpc/go/sonr-io/motor/api/v1"
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
		bind.Instance, err = mtr.NewWasmMotor(req.DeviceId)
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

type tuple struct {
	Result []byte
	Error  error
}

// const [res, error] = await init();
func InitExporter() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) tuple {
		b, err := json.Marshal(args[0])
		if err != nil {
			return tuple{
				Result: nil,
				Error:  err,
			}
		}
		res, err := Init(b)
		if err != nil {
			return tuple{
				Result: nil,
				Error:  err,
			}
		}

		return tuple{
			Result: res,
			Error:  nil,
		}
	})
}

func CreateAccountExporter() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) tuple {
		b, err := json.Marshal(args[0])
		if err != nil {
			return tuple{
				Result: nil,
				Error:  err,
			}
		}

		res, err := bind.CreateAccount(b)
		return tuple{
			Result: res,
			Error:  err,
		}
	})
}

func LoginExporter() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) tuple {
		b, err := json.Marshal(args[0])
		if err != nil {
			return tuple{
				Result: nil,
				Error:  err,
			}
		}

		res, err := bind.Login(b)
		return tuple{
			Result: res,
			Error:  err,
		}
	})
}

func AddressExporter() js.Func {
	js_func := js.FuncOf(func(this js.Value, args []js.Value) []byte {
		return []byte(bind.Address())
	})

	return js_func
}

func BalanceExporter() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) int {
		return bind.Balance()
	})
}

func DidDocExporter() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) tuple {
		res, err := bind.DidDoc()
		return tuple{
			Result: res,
			Error:  err,
		}
	})
}

func main() {
	js.Global().Set("init", InitExporter())
	js.Global().Set("createAccount", CreateAccountExporter())
	js.Global().Set("login", LoginExporter())
	js.Global().Set("getAddress", AddressExporter())
	js.Global().Set("getBalance", BalanceExporter())
	js.Global().Set("getDidDocument", DidDocExporter())

	// module cannot leave scope, keep the entry in scope
	<-make(chan (bool))
}
