package keeper

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/sonr-io/sonr/x/schema/keeper"
	"github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func SchemaKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"SchemaParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		nil,
		nil,
		nil,
		nil,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}

func CreateWhatIsWithDID(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.WhatIs {
	items := make([]types.WhatIs, n)
	var offset int = n + 1
	for i := range items {
		id := fmt.Sprintf("did:snr:%s", strconv.Itoa(i))
		schemaId := fmt.Sprintf("did:snr:%s", strconv.Itoa(offset))
		var whatIs = types.WhatIs{
			Creator: id,
			Did:     schemaId,
			Schema: &types.SchemaDefinition{
				Did:   schemaId,
				Label: "test",
				Fields: []*types.SchemaKindDefinition{
					{
						Name:  "name",
						Field: types.SchemaKind_STRING,
					},
					{
						Name:  "age",
						Field: types.SchemaKind_INT,
					},
					{
						Name:  "DOB",
						Field: types.SchemaKind_STRING,
					},
				},
			},
			Timestamp: time.Now().Unix(),
			IsActive:  true,
		}
		keeper.SetWhatIs(ctx, whatIs)
		items[i] = whatIs
		offset += 1
	}
	return items
}
