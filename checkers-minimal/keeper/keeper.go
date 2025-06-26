package keeper

import (
	"fmt"
	"github.com/alice/checkers/rules"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/tendermint/tendermint/libs/log"

	storetype "cosmossdk.io/store/types"
	"github.com/alice/checkers"
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		bank       rules.BankEscrowKeeper
		hooks      rules.CheckersHooks
		cdc        codec.BinaryCodec
		storeKey   StoreKey
		memKey     StoreKey
		paramstore paramtypes.Subspace
	}
)

type StoreKey interface {
	Name() string
	String() string
}

// NewKeeper creates a new Keeper instance
func NewKeeper(
	bank rules.BankEscrowKeeper,
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetype.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(rules.ParamKeyTable())
	}

	return &Keeper{
		bank:       bank,
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}
