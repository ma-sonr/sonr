//go:build wasm
// +build wasm

package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/sonr-io/sonr/bind"
)

type tuple struct {
	Result []byte
	Error  error
}

func InitExporter() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b, err := json.Marshal(args[0])
		if err != nil {
			return tuple{
				Result: nil,
				Error:  err,
			}
		}
		res, err := bind.Init(b)
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
