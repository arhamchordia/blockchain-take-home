package keeper

import (
	"context"
	"strconv"

	"blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreatePost creates a new blog post
func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	currentTime := ctx.BlockHeader().Time

	post := types.Post{
		Creator:       msg.Creator,
		Title:         msg.Title,
		Body:          msg.Body,
		CreatedAt:     currentTime,
		LastUpdatedAt: currentTime,
		Editors:       []string{msg.Creator},
	}

	id := k.AppendPost(ctx, post)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreatePost,
			sdk.NewAttribute(types.AttributeKeyPostID, strconv.FormatUint(id, 10)),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyTitle, msg.Title),
		),
	)

	return &types.MsgCreatePostResponse{
		Id: id,
	}, nil
}
