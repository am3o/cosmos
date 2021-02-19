package cosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/am3o/cosmos/x/cosmos/keeper"
	"github.com/am3o/cosmos/x/cosmos/types"
)

// Handle a message to delete name
func handleMsgDeleteVote(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteVote) (*sdk.Result, error) {
	if !k.VoteExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetVoteOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteVote(ctx, msg.ID)
	return &sdk.Result{}, nil
}