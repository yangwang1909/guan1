package keeper

import (
	"context"

	"voter/x/tokenfactory/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) MintAndSendTokens(ctx context.Context, msg *types.MsgMintAndSendTokens) (*types.MsgMintAndSendTokensResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Owner); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgMintAndSendTokensResponse{}, nil
}
