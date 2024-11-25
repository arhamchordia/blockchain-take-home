package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddEditor{}

func NewMsgAddEditor(creator string, id uint64, editor string) *MsgAddEditor {
	return &MsgAddEditor{
		Creator: creator,
		Id:      id,
		Editor:  editor,
	}
}

func (msg *MsgAddEditor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Editor)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid editor address (%s)", err)
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteEditor{}

func NewMsgDeleteEditor(creator string, id uint64, editor string) *MsgDeleteEditor {
	return &MsgDeleteEditor{
		Creator: creator,
		Id:      id,
		Editor:  editor,
	}
}

func (msg *MsgDeleteEditor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Editor)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid editor address (%s)", err)
	}

	return nil
}
