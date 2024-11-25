package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"blog/x/blog/types"
)

func (k msgServer) UpdatePost(goCtx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	_, foundEditor := k.checkIfExists(val, msg.Creator)
	if !foundEditor {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect editor")
	}

	// update val details
	val.LastUpdatedAt = ctx.BlockHeader().Time
	val.Body = msg.Body
	val.Title = msg.Title

	k.SetPost(ctx, val)
	return &types.MsgUpdatePostResponse{}, nil
}
