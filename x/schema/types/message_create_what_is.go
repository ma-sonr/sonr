package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateWhatIs = "create_what_is"

var _ sdk.Msg = &MsgCreateWhatIs{}

func NewMsgCreateWhatIs(creator string) *MsgCreateWhatIs {
  return &MsgCreateWhatIs{
		Creator: creator,
	}
}

func (msg *MsgCreateWhatIs) Route() string {
  return RouterKey
}

func (msg *MsgCreateWhatIs) Type() string {
  return TypeMsgCreateWhatIs
}

func (msg *MsgCreateWhatIs) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateWhatIs) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWhatIs) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

