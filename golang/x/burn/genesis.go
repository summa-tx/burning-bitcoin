package burn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// GenesisState is the genesis state
type GenesisState struct {
}

// NewGenesisState instantiates a genesis state
func NewGenesisState() GenesisState {
	return GenesisState{}
}

// ValidateGenesis Validates a genesis state
func ValidateGenesis(data GenesisState) error {
	return nil
}

// DefaultGenesisState makes a default empty genesis state
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// InitGenesis inits the app state based on the genesis state
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// ExportGenesis exports the genesis state
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	panic("Not implemented")
}
