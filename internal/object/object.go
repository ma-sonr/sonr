package object

import (
	"errors"

	"github.com/sonr-io/sonr/internal/schemas"
	"github.com/sonr-io/sonr/pkg/protocol"
)

var (
	ErrObjectNotUploaded = errors.New("error while uploading object")
	ErrObjectEmpty       = errors.New("object cannot be empty")
)

type ObjectReference struct {
	Did   string
	Label string
	Cid   string
}

/*
	Object definition to be returned after object creation
*/
type ObjectUploadResult struct {
	Code      int32
	Reference *ObjectReference
	Message   string
}

type objectImpl struct {
	store  protocol.IPFS
	schema schemas.Schema
}

func New(schemaImpl schemas.Schema, store protocol.IPFS) *objectImpl {
	return &objectImpl{
		// TODO: replace with store interface that Daniel made
		store:  store,
		schema: schemaImpl,
	}
}

func NewWithConfig(c *Config) *objectImpl {
	return &objectImpl{
		store:  c.storeImpl,
		schema: c.schema,
	}
}
