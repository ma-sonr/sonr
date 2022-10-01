package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sonr-io/sonr/x/schema/types"
)

var _ = strconv.Itoa(0)

func CmdCreateWhatIs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-what-is",
		Short: "Broadcast message CreateWhatIs",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateWhatIs(
				clientCtx.GetFromAddress().String(),
				
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}