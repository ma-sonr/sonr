package schemas

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/sonr-io/sonr/pkg/protocol"
	"github.com/sonr-io/sonr/pkg/store"
	st "github.com/sonr-io/sonr/x/schema/types"
)

var (
	errSchemaFieldsInvalid  = errors.New("supplied Schema is invalid")
	errSchemaFieldsNotFound = errors.New("no Schema Fields found")
	errNodeNotFound         = errors.New("no object definition built from schema")
)

type Encoding int

type Event struct {
	name     string
	previous cid.Cid
}

const (
	EncType_DAG_CBOR Encoding = iota
	EncType_DAG_JSON
)

type schemaImpl struct {
	fields     []*st.SchemaKindDefinition
	subSchemas map[string]*st.SchemaDefinition
	whatIs     *st.WhatIs
	nodes      datamodel.Node
	store      protocol.IPFS
	next       *schemaImpl
}

/*
	Default initialization with a local shell for persistence
*/
func New(fields []*st.SchemaKindDefinition, whatIs *st.WhatIs) Schema {
	asi := &schemaImpl{
		fields:     fields,
		subSchemas: make(map[string]*st.SchemaDefinition),
		whatIs:     whatIs,
		nodes:      nil,
		store:      protocol.NewIPFSShell("localhost:5001", store.NewMemoryStore().Batching()),
	}

	asi.loadSubSchemas(context.Background(), fields)
	return asi
}

/*
	Initialize with a ipfs shell instance
*/
func NewWithShell(store protocol.IPFS, fields []*st.SchemaKindDefinition, whatIs *st.WhatIs) Schema {
	asi := &schemaImpl{
		fields:     fields,
		subSchemas: make(map[string]*st.SchemaDefinition),
		whatIs:     whatIs,
		nodes:      nil,
		store:      store,
	}

	asi.loadSubSchemas(context.Background(), fields)
	return asi
}
