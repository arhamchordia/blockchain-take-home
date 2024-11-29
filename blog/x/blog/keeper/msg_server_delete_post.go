package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"blog/x/blog/types"
)

// DeletePost deletes an existing blog post if the creator has proper authorization
func (k msgServer) DeletePost(goCtx context.Context, msg *types.MsgDeletePost) (*types.MsgDeletePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate basic message properties
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the post
	post, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Check authorization
	if !k.HasEditor(post, msg.Creator) {
		return nil, errorsmod.Wrapf(
			sdkerrors.ErrUnauthorized,
			"address %s is not authorized to delete post %d",
			msg.Creator,
			msg.Id,
		)
	}

	// Remove the post
	k.RemovePost(ctx, msg.Id)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDeletePost,
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.Id)),
			sdk.NewAttribute(types.AttributeKeyDeleter, msg.Creator),
		),
	)

	return &types.MsgDeletePostResponse{}, nil
}
