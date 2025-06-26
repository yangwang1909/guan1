package keeper

import (
	"context"
	"github.com/alice/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PlayMove(goCtx context.Context, msg *rules.MsgPlayMove) (*rules.MsgPlayMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &rules.MsgPlayMoveResponse{}, nil
}
