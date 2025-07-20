package keeper

import (
	"context"

	"voter/x/tokenfactory/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) UpdateOwner(ctx context.Context, msg *types.MsgUpdateOwner) (*types.MsgUpdateOwnerResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Owner); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgUpdateOwnerResponse{}, nil
}
