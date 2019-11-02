package keeper

import (
	"bytes"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/summa-tx/bitcoin-spv/golang/btcspv"
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

func (k Keeper) setValidated(ctx sdk.Context, proof btcspv.SPVProof) {
	store := ctx.KVStore(k.storeKey)

	enc := new(bytes.Buffer)
	err := binary.Write(enc, binary.LittleEndian, proof.ConfirmingHeader.Height)
	if err != nil {
		panic("")
	}

	store.Set(proof.TxID[:], enc.Bytes())
}

// SetValidated sets a thing
func (k Keeper) SetValidated(ctx sdk.Context, proof btcspv.SPVProof) {
	k.setValidated(ctx, proof)
}

// GetValidated gets a thing
func (k Keeper) GetValidated(ctx sdk.Context, txid [32]byte) uint {
	store := ctx.KVStore(k.storeKey)
	return btcspv.BytesToUint(store.Get(txid[:]))
}

// AppendAddr appends a new ponzi burner
func (k Keeper) AppendAddr(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	kings := store.Get([]byte("ponzikings"))

	kings = append(kings, addr.Bytes()...)

	store.Set([]byte("ponzikings"), kings)
}

// GetKings gets the list of kings
func (k Keeper) GetKings(ctx sdk.Context) []sdk.AccAddress {
	store := ctx.KVStore(k.storeKey)
	kings := store.Get([]byte("ponzikings"))

	var kingsAddrs []sdk.AccAddress
	for i := 0; i < len(kings); i += 20 {
		kingsAddrs = append(kingsAddrs, sdk.AccAddress(kings[i:i+20]))
	}

	return kingsAddrs
}
