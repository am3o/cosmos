package cosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/am3o/cosmos/x/cosmos/keeper"
	"github.com/am3o/cosmos/x/cosmos/types"
)

func handleMsgCreateVote(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateVote) (*sdk.Result, error) {
	k.CreateVote(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
