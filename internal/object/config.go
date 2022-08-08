package object

import (
	"github.com/sonr-io/sonr/internal/schemas"
	"github.com/sonr-io/sonr/pkg/protocol"
)

type Config struct {
	storeImpl protocol.IPFS
	schema    schemas.Schema
}

func (c *Config) WithStorage(store protocol.IPFS) {
	c.storeImpl = store
}

func (c *Config) WithSchemaImpl(schema schemas.Schema) {
	c.schema = schema
}

type ObjectConfiguration = func(config *Config)
