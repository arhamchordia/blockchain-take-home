package keeper_test

import (
	"context"
	"fmt"
	"github.com/cometbft/cometbft/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"

	keepertest "blog/testutil/keeper"
	"blog/x/blog/keeper"
	"blog/x/blog/types"
)

var creator1 sdk.AccAddress = sdk.AccAddress(crypto.AddressHash([]byte("creator")))
var creator2 sdk.AccAddress = sdk.AccAddress(crypto.AddressHash([]byte("editor")))

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.BlogKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}

func TestMsgCreatePost(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	fmt.Println(creator1.String(), creator2.String())
	fmt.Println(k.GetPostCount(wctx))

	testCases := []struct {
		name      string
		input     *types.MsgCreatePost
		expErr    bool
		expErrMsg string
	}{
		{
			name: "valid post creation",
			input: &types.MsgCreatePost{
				Creator: creator1.String(),
				Title:   "Test Title",
				Body:    "This is the body of the post",
			},
			expErr: false,
		},
		{
			name: "missing title",
			input: &types.MsgCreatePost{
				Creator: creator1.String(),
				Title:   "",
				Body:    "This is the body of the post",
			},
			expErr:    true,
			expErrMsg: "missing title: invalid request",
		},
		{
			name: "missing creator",
			input: &types.MsgCreatePost{
				Creator: "abcd",
				Title:   "Test Title",
				Body:    "This is the body of the post",
			},
			expErr:    true,
			expErrMsg: "invalid creator address (decoding bech32 failed: invalid bech32 string length 4): invalid address",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.CreatePost(wctx, tc.input)
			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgDeletePost(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	post := types.Post{
		Id:      1,
		Creator: creator1.String(),
		Title:   "Existing Post 1",
		Body:    "This is a pre-existing post. 1",
		Editors: []string{creator1.String()},
	}
	k.SetPost(wctx, post)

	post = types.Post{
		Id:      2,
		Creator: creator1.String(),
		Title:   "Existing Post 2",
		Body:    "This is a pre-existing post. 2",
		Editors: []string{creator1.String()},
	}
	k.SetPost(wctx, post)

	testCases := []struct {
		name      string
		input     *types.MsgDeletePost
		expErr    bool
		expErrMsg string
	}{
		{
			name: "valid post deletion",
			input: &types.MsgDeletePost{
				Id:      1,
				Creator: creator1.String(),
			},
			expErr: false,
		},
		{
			name: "post not found",
			input: &types.MsgDeletePost{
				Id:      999,
				Creator: creator1.String(),
			},
			expErr:    true,
			expErrMsg: "key 999 doesn't exist",
		},
		{
			name: "unauthorized deletion",
			input: &types.MsgDeletePost{
				Id:      2,
				Creator: creator2.String(),
			},
			expErr:    true,
			expErrMsg: "incorrect editor: unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.DeletePost(wctx, tc.input)

			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrMsg)
			} else {
				require.NoError(t, err)

				_, found := k.GetPost(wctx, tc.input.Id)
				require.False(t, found)
			}
		})
	}
}

func TestMsgUpdatePost(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	originalPost := types.Post{
		Id:            1,
		Creator:       creator1.String(),
		Title:         "Original Title",
		Body:          "Original Body",
		CreatedAt:     wctx.BlockHeader().Time,
		LastUpdatedAt: wctx.BlockHeader().Time,
		Editors:       []string{creator1.String()},
	}
	k.SetPost(wctx, originalPost)

	testCases := []struct {
		name      string
		input     *types.MsgUpdatePost
		expErr    bool
		expErrMsg string
	}{
		{
			name: "valid update",
			input: &types.MsgUpdatePost{
				Id:      1,
				Creator: creator1.String(),
				Title:   "Updated Title",
				Body:    "Updated Body",
			},
			expErr: false,
		},
		{
			name: "post not found",
			input: &types.MsgUpdatePost{
				Id:      999,
				Creator: creator1.String(),
				Title:   "Doesn't Matter",
				Body:    "Doesn't Matter",
			},
			expErr:    true,
			expErrMsg: "key 999 doesn't exist",
		},
		{
			name: "unauthorized update",
			input: &types.MsgUpdatePost{
				Id:      1,
				Creator: creator2.String(),
				Title:   "Hacked Title",
				Body:    "Hacked Body",
			},
			expErr:    true,
			expErrMsg: "incorrect editor: unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.UpdatePost(wctx, tc.input)

			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrMsg)
			} else {
				require.NoError(t, err)

				// Verify that the post was updated
				updatedPost, found := k.GetPost(wctx, tc.input.Id)
				require.True(t, found, "Post should exist after update")
				require.Equal(t, tc.input.Title, updatedPost.Title, "Title should be updated")
				require.Equal(t, tc.input.Body, updatedPost.Body, "Body should be updated")
			}
		})
	}
}

func TestAddEditor(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	// Setup: create a post
	post := types.Post{
		Id:      1,
		Creator: creator1.String(),
		Title:   "Sample Post",
		Body:    "Sample Body",
		Editors: []string{creator1.String()},
	}
	k.SetPost(wctx, post)

	// Test: Add a new editor
	msg := &types.MsgAddEditor{
		Creator: creator1.String(),
		Id:      1,
		Editor:  creator2.String(),
	}
	_, err := ms.AddEditor(wctx, msg)
	require.NoError(t, err, "Adding a valid editor should not return an error")

	// Verify: The editor was added
	updatedPost, found := k.GetPost(wctx, 1)
	require.True(t, found, "Post should exist")
	require.Contains(t, updatedPost.Editors, creator2.String(), "Editor should be added to the post")

	// Test: Add an existing editor
	msg.Editor = creator2.String()
	_, err = ms.AddEditor(wctx, msg)
	require.Error(t, err, "Adding an existing editor should return an error")

	// Test: Unauthorized access
	msg.Creator = creator2.String()
	_, err = ms.AddEditor(wctx, msg)
	require.Error(t, err, "Adding an editor by an unauthorized user should return an error")
}

func TestDeleteEditor(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	// Setup: create a post
	post := types.Post{
		Id:      1,
		Creator: creator1.String(),
		Title:   "Sample Post",
		Body:    "Sample Body",
		Editors: []string{creator1.String(), creator2.String()},
	}
	k.SetPost(wctx, post)

	// Test: Delete an existing editor
	msg := &types.MsgDeleteEditor{
		Creator: creator1.String(),
		Id:      1,
		Editor:  creator2.String(),
	}
	_, err := ms.DeleteEditor(wctx, msg)
	require.NoError(t, err, "Deleting a valid editor should not return an error")

	// Verify: The editor was removed
	updatedPost, found := k.GetPost(wctx, 1)
	require.True(t, found, "Post should exist")
	require.NotContains(t, updatedPost.Editors, creator2.String(), "Editor should be removed from the post")

	// Test: Delete a non-existent editor
	msg.Editor = creator2.String()
	_, err = ms.DeleteEditor(wctx, msg)
	require.EqualError(t, err, "editor does not exist: invalid address")

	// Test: Unauthorized access
	msg.Creator = creator2.String()
	msg.Editor = creator2.String()
	_, err = ms.DeleteEditor(wctx, msg)
	require.EqualError(t, err, "incorrect owner: unauthorized")

	// Test: Creator cannot be deleted as editor
	msg.Creator = creator1.String()
	msg.Editor = creator1.String()
	_, err = ms.DeleteEditor(wctx, msg)
	require.EqualError(t, err, "creator cannot be deleted from editors: unauthorized")
}
