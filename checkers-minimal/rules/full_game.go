package rules

import (
	"fmt"
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (storedGame StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, errBlack := sdk.AccAddressFromBech32(storedGame.Black)
	return black, sdkerrors.Wrapf(errBlack, ErrInvalidBlack.Error(), storedGame.Black)
}

func (storedGame StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
	red, errRed := sdk.AccAddressFromBech32(storedGame.Red)
	return red, sdkerrors.Wrapf(errRed, ErrInvalidRed.Error(), storedGame.Red)
}

func (storedGame StoredGame) ParseGame() (game *Game, err error) {
	board, errBoard := Parse(storedGame.Board)
	if errBoard != nil {
		return nil, sdkerrors.Wrapf(errBoard, ErrGameNotParseable.Error())
	}
	board.Turn = StringPieces[storedGame.Turn].Player
	if board.Turn.Color == "" {
		return nil, sdkerrors.Wrapf(fmt.Errorf("turn: %s", storedGame.Turn), ErrGameNotParseable.Error())
	}
	return board, nil
}

func (storedGame StoredGame) GetPlayerAddress(color string) (address sdk.AccAddress, found bool, err error) {
	black, err := storedGame.GetBlackAddress()
	if err != nil {
		return nil, false, err
	}
	red, err := storedGame.GetRedAddress()
	if err != nil {
		return nil, false, err
	}
	address, found = map[string]sdk.AccAddress{
		PieceStrings[BLACK_PLAYER]: black,
		PieceStrings[RED_PLAYER]:   red,
	}[color]
	return address, found, nil
}

func (storedGame StoredGame) GetWinnerAddress() (address sdk.AccAddress, found bool, err error) {
	return storedGame.GetPlayerAddress(storedGame.Winner)
}

func (storedGame StoredGame) GetDeadlineAsTime() (deadline time.Time, err error) {
	deadline, errDeadline := time.Parse("2006-01-02 15:04:05.999999999 +0000 UTC", storedGame.Deadline)
	return deadline, sdkerrors.Wrapf(errDeadline, ErrInvalidDeadline.Error(), storedGame.Deadline)
}

func (storedGame StoredGame) Validate() (err error) {
	_, err = storedGame.GetBlackAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.GetRedAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.ParseGame()
	if err != nil {
		return err
	}
	_, err = storedGame.GetDeadlineAsTime()
	return err
}
