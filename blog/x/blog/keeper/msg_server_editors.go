package keeper

import (
	"blog/x/blog/types"
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// AddEditor adds a new editor to a post if all validations pass
func (k msgServer) AddEditor(goCtx context.Context, msg *types.MsgAddEditor) (*types.MsgAddEditorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	post, err := k.validatePostAndOwnership(ctx, msg.Id, msg.Creator)
	if err != nil {
		return nil, err
	}

	if k.HasEditor(post, msg.Editor) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "editor already exists")
	}

	post.Editors = append(post.Editors, msg.Editor)
	k.SetPost(ctx, post)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddEditor,
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.Id)),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyEditor, msg.Editor),
		),
	)

	return &types.MsgAddEditorResponse{}, nil
}

// DeleteEditor removes an editor from a post if all validations pass
func (k msgServer) DeleteEditor(goCtx context.Context, msg *types.MsgDeleteEditor) (*types.MsgDeleteEditorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	post, err := k.validatePostAndOwnership(ctx, msg.Id, msg.Creator)
	if err != nil {
		return nil, err
	}

	if msg.Editor == post.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "creator cannot be deleted from editors")
	}

	editorIndex, found := k.FindEditorIndex(post, msg.Editor)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "editor already exists")
	}

	post.Editors = append(post.Editors[:editorIndex], post.Editors[editorIndex+1:]...)
	k.SetPost(ctx, post)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDeleteEditor,
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.Id)),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyEditor, msg.Editor),
		),
	)

	return &types.MsgDeleteEditorResponse{}, nil
}

// validatePostAndOwnership checks if post exists and if the creator is the owner
func (k msgServer) validatePostAndOwnership(ctx sdk.Context, postID uint64, creator string) (types.Post, error) {
	post, found := k.GetPost(ctx, postID)
	if !found {
		return types.Post{}, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", postID))
	}

	if creator != post.Creator {
		return types.Post{}, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	return post, nil
}

// HasEditor checks if an address is already an editor
func (k msgServer) HasEditor(post types.Post, address string) bool {
	_, found := k.FindEditorIndex(post, address)
	return found
}

// FindEditorIndex returns the index of an editor in the editors list
func (k msgServer) FindEditorIndex(post types.Post, address string) (int, bool) {
	for i, editor := range post.Editors {
		if editor == address {
			return i, true
		}
	}
	return -1, false
}
