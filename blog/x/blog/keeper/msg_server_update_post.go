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

	foundEditor := k.hasEditor(val, msg.Creator)
	if !foundEditor {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect editor")
	}

	// update val details
	val.LastUpdatedAt = ctx.BlockHeader().Time
	val.Body = msg.Body
	val.Title = msg.Title
	k.SetPost(ctx, val)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdatePost,
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.Id)),
			sdk.NewAttribute(types.AttributeKeyEditor, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyTitle, msg.Title),
			sdk.NewAttribute(types.AttributeKeyUpdateTime, val.LastUpdatedAt.String()),
		),
	)

	return &types.MsgUpdatePostResponse{}, nil
}
