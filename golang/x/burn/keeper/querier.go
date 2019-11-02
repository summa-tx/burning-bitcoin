package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/summa-tx/burning-bitcoin/golang/x/burn/types"
)

// QuerySomeThing is a query string tag for get something
const QuerySomeThing = "something"

// NewQuerier makes a query routing function
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QuerySomeThing:
			return querySomeThing(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown query endpoint")
		}
	}
}

func querySomeThing(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResSomeThing{})

	if err != nil {
		// TODO: make more good by returning sdk error
		panic("could not marshal result to JSON")
	}

	return res, nil
}
