package keeper

import (
	"blog/x/blog/types"
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddEditor(goCtx context.Context, msg *types.MsgAddEditor) (*types.MsgAddEditorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	_, found = k.checkIfExists(val, msg.Editor)
	if found {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "editor already exists")
	}

	// set if all checks pass
	val.Editors = append(val.Editors, msg.Editor)
	k.SetPost(ctx, val)

	return &types.MsgAddEditorResponse{}, nil
}

func (k msgServer) DeleteEditor(goCtx context.Context, msg *types.MsgDeleteEditor) (*types.MsgDeleteEditorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	i, found := k.checkIfExists(val, msg.Editor)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "editor does not exist")
	}

	if msg.Editor == val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "creator cannot be deleted from editors")
	}

	// set if all checks pass
	val.Editors = append(val.Editors[:i], val.Editors[i+1:]...)
	k.SetPost(ctx, val)

	return &types.MsgDeleteEditorResponse{}, nil
}

func (k Keeper) checkIfExists(post types.Post, address string) (int, bool) {
	for i, editor := range post.Editors {
		if editor == address {
			return i, true
		}
	}
	return 0, false
}
