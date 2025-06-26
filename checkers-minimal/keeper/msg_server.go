package keeper

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"cosmossdk.io/collections"
	"github.com/alice/checkers"
	"github.com/alice/checkers/rules"
)

type msgServer struct {
	k Keeper
}

var _ checkers.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) checkers.MsgServer {
	return &msgServer{k: keeper}
}

// CreateGame defines the handler for the MsgCreateGame message.
func (ms msgServer) CreateGame(ctx context.Context, msg *checkers.MsgCreateGame) (*checkers.MsgCreateGameResponse, error) {
	if length := len([]byte(msg.Index)); checkers.MaxIndexLength < length || length < 1 {
		return nil, checkers.ErrIndexTooLong
	}
	if _, err := ms.k.StoredGames.Get(ctx, msg.Index); err == nil || errors.Is(err, collections.ErrEncoding) {
		return nil, fmt.Errorf("game already exists at index: %s", msg.Index)
	}
	systemInfo, found := ms.k.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)

	newGame := rules.New()
	storedGame := rules.StoredGame{
		Index: newIndex,
		Board: newGame.String(),
		Turn:  rules.PieceStrings[newGame.Turn],
		Black: msg.Black,
		Red:   msg.Red,
		Wager: msg.Wager,
	}

	if err := storedGame.Validate(); err != nil {
		return nil, err
	}
	if err := ms.k.StoredGames.Set(ctx, msg.Index, storedGame); err != nil {
		return nil, err
	}

	return &checkers.MsgCreateGameResponse{}, nil
}
