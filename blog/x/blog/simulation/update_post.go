package simulation

import (
	"fmt"
	"math/rand"
	"slices"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"blog/x/blog/keeper"
	"blog/x/blog/types"
)

func SimulateMsgUpdatePost(
	txGen client.TxConfig,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msgType := sdk.MsgTypeURL(&types.MsgUpdatePost{})

		simAccount, _ := simtypes.RandomAcc(r, accs)

		fees, err := simtypes.RandomFees(r, ctx, bk.SpendableCoins(ctx, ak.GetAccount(ctx, simAccount.Address).GetAddress()))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msgType, "unable to generate fees"), nil, err
		}

		postCount := k.GetPostCount(ctx)
		if postCount == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msgType, "no posts in there yet"), nil, nil
		}

		postId := uint64(RandRange(1, int(postCount+1)))
		post, err := k.ShowPost(ctx, &types.QueryShowPostRequest{Id: postId})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msgType, "unable to find post"), nil, nil
		}

		if !slices.Contains(post.Post.Editors, ak.GetAccount(ctx, simAccount.Address).GetAddress().String()) {
			err = errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not authorized to update post %d", simAccount.Address, postId)
			return simtypes.NoOpMsg(types.ModuleName, msgType, "post editor does not exist"), nil, nil
		}

		msg := types.NewMsgUpdatePost(simAccount.Address.String(), fmt.Sprintf("title-%d", r.Intn(1000)), fmt.Sprintf("title-%d", r.Intn(1000)), postId)

		txCtx := simulation.OperationInput{
			R:             r,
			App:           app,
			TxGen:         txGen,
			Cdc:           nil,
			Msg:           msg,
			Context:       ctx,
			SimAccount:    simAccount,
			AccountKeeper: ak,
			ModuleName:    types.ModuleName,
		}

		return simulation.GenAndDeliverTx(txCtx, fees)
	}
}
