package functions

import (
	"strings"
)

type FunctionInterface interface {
	Store(f *Function) (string, error)
	GetAndExecute(path string, writer *strings.Builder) error
	Execute(function *Function, writer *strings.Builder) error
}
