package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"blog/x/blog/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check message
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	post := types.Post{
		Creator:       msg.Creator,
		Title:         msg.Title,
		Body:          msg.Body,
		CreatedAt:     ctx.BlockHeader().Time,
		LastUpdatedAt: ctx.BlockHeader().Time,
		Editors:       []string{msg.Creator},
	}
	id := k.AppendPost(
		ctx,
		post,
	)
	return &types.MsgCreatePostResponse{
		Id: id,
	}, nil
}
