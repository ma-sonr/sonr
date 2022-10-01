package simulation

import (
	"math/rand"

	"github.com/sonr-io/sonr/x/schema/keeper"
	"github.com/sonr-io/sonr/x/schema/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateWhatIs(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateWhatIs{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateWhatIs simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateWhatIs simulation not implemented"), nil, nil
	}
}
