package keeper

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper instantiates a new keeper
func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

func (k Keeper) setValidated(ctx sdk.Context, txid [32]byte) {
	store := ctx.KVStore(k.storeKey)
	store.Set(txid[:], []byte{1})
}

// SetValidated sets a thing
func (k Keeper) SetValidated(ctx sdk.Context, txid [32]byte) {
	k.setValidated(ctx, txid)
}

// GetValidated gets a thing
func (k Keeper) GetValidated(ctx sdk.Context, txid [32]byte) bool {
	store := ctx.KVStore(k.storeKey)
	return bytes.Equal(store.Get(txid[:]), []byte{1})
}
