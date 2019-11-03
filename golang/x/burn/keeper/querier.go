package keeper

import (
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/summa-tx/burning-bitcoin/golang/x/burn/types"
)

// QueryValidated is a query string tag for get something
const QueryValidated = "validated"

// NewQuerier makes a query routing function
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryValidated:
			return queryValidated(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown query endpoint")
		}
	}
}

func queryValidated(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	digest, err := hex.DecodeString(path[0])
	if err != nil || len(digest) != 32 {
		// TODO: make more good by returning sdk error
		panic("bad")
	}

	var fixedDigest [32]byte
	copy(fixedDigest[:], digest)

	txid := hex.EncodeToString(fixedDigest[:])
	validated := keeper.GetValidated(ctx, fixedDigest)
	res, err := codec.MarshalJSONIndent(
		keeper.cdc,
		types.QueryResValidated{
			TxID:   txid,
			Result: validated})

	if err != nil {
		// TODO: make more good by returning sdk error
		panic("could not marshal result to JSON")
	}

	return res, nil
}
