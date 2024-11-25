package keeper_test

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"

	keepertest "blog/testutil/keeper"
	"blog/x/blog/keeper"
	"blog/x/blog/types"
)

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
				Creator: "cosmos1ly902zr06vgg6kkqhetnvka4rl66qza795xac4",
				Title:   "Test Title",
				Body:    "This is the body of the post",
			},
			expErr: false,
		},
		{
			name: "missing title",
			input: &types.MsgCreatePost{
				Creator: "cosmos1ly902zr06vgg6kkqhetnvka4rl66qza795xac4",
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
		Creator: "cosmos1creator...",
		Title:   "Existing Post 1",
		Body:    "This is a pre-existing post. 1",
	}
	k.SetPost(wctx, post)

	post = types.Post{
		Id:      2,
		Creator: "cosmos1creator...",
		Title:   "Existing Post 2",
		Body:    "This is a pre-existing post. 2",
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
				Creator: "cosmos1creator...",
			},
			expErr: false,
		},
		{
			name: "post not found",
			input: &types.MsgDeletePost{
				Id:      999,
				Creator: "cosmos1creator...",
			},
			expErr:    true,
			expErrMsg: "key 999 doesn't exist",
		},
		{
			name: "unauthorized deletion",
			input: &types.MsgDeletePost{
				Id:      2,
				Creator: "cosmos1unauthorized...",
			},
			expErr:    true,
			expErrMsg: "incorrect owner",
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
		Creator:       "cosmos1ly902zr06vgg6kkqhetnvka4rl66qza795xac4",
		Title:         "Original Title",
		Body:          "Original Body",
		CreatedAt:     wctx.BlockHeader().Time,
		LastUpdatedAt: wctx.BlockHeader().Time,
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
				Creator: "cosmos1ly902zr06vgg6kkqhetnvka4rl66qza795xac4",
				Title:   "Updated Title",
				Body:    "Updated Body",
			},
			expErr: false,
		},
		{
			name: "post not found",
			input: &types.MsgUpdatePost{
				Id:      999,
				Creator: "cosmos1ly902zr06vgg6kkqhetnvka4rl66qza795xac4",
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
				Creator: "cosmos1twuaz9rejs62uj89zexvnvwgj5mh0mntvnmuas",
				Title:   "Hacked Title",
				Body:    "Hacked Body",
			},
			expErr:    true,
			expErrMsg: "incorrect owner",
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
