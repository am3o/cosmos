package cosmos

import (
	"fmt"

	"github.com/am3o/cosmos/x/cosmos/keeper"
	"github.com/am3o/cosmos/x/cosmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

func handleMsgCreateVote(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateVote) (*sdk.Result, error) {
	k.CreateVote(ctx, msg)

	coin, _ := sdk.ParseCoins("50token")

	_, err := k.CoinKeeper.AddCoins(ctx, msg.Creator.Bytes(), coin)
	if err != nil {
		return nil, err
	}

	var accAddress = crypto.AddressHash([]byte(types.ModuleName))
	moduleAcct := sdk.AccAddress(accAddress)

	k.CoinKeeper.SubtractCoins(ctx, moduleAcct, coin)
	if err != nil {
		return nil, err
	}
	
	fmt.Println(k.CoinKeeper.GetCoins(ctx, moduleAcct))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
