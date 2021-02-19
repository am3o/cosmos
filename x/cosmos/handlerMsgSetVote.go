package cosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/am3o/cosmos/x/cosmos/keeper"
	"github.com/am3o/cosmos/x/cosmos/types"
)

func handleMsgSetVote(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetVote) (*sdk.Result, error) {
	var vote = types.Vote{
		Creator: msg.Creator,
		ID:      msg.ID,
		PollID:  msg.PollID,
		Value:   msg.Value,
	}
	if !msg.Creator.Equals(k.GetVoteOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetVote(ctx, vote)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
