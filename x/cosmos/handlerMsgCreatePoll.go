package cosmos

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	"github.com/am3o/cosmos/x/cosmos/keeper"
	"github.com/am3o/cosmos/x/cosmos/types"
)

func handleMsgCreatePoll(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreatePoll) (*sdk.Result, error) {
	var poll = types.Poll{
		Creator: msg.Creator,
		Title:   msg.Title,
		Options: msg.Options,
	}

	var accAddress = crypto.AddressHash([]byte(types.ModuleName))
	fmt.Println(accAddress)

	moduleAcct := sdk.AccAddress(accAddress)
	payment, _ := sdk.ParseCoins("200token")
	if err := k.CoinKeeper.SendCoins(ctx, poll.Creator, moduleAcct, payment); err != nil {
		return nil, err
	}
	k.CreatePoll(ctx, msg)

	fmt.Println(k.CoinKeeper.GetCoins(ctx, moduleAcct))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
