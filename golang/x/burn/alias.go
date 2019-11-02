package burn

import (
	"github.com/summa-tx/burning-bitcoin/golang/x/burn/keeper"
	"github.com/summa-tx/burning-bitcoin/golang/x/burn/types"
)

const (
	// ModuleName is what it says on the tin
	ModuleName = types.ModuleName
	// RouterKey is what it says on the tin
	RouterKey = types.RouterKey
	//StoreKey is what it says on the tin
	StoreKey = types.StoreKey
)

var (
	// NewKeeper is what is says on the tin
	NewKeeper = keeper.NewKeeper
	// NewQuerier is what is says on the tin
	NewQuerier = keeper.NewQuerier
	// NewMsgBurnProof is what is says on the tin
	NewMsgBurnProof = types.NewMsgBurnProof
	// RegisterCodec is what is says on the tin
	RegisterCodec = types.RegisterCodec
	// ModuleCdc is what is says on the tin
	ModuleCdc = types.ModuleCdc
)

type (
	// Keeper is what is says on the tin
	Keeper = keeper.Keeper
	// MsgBurnProof is what is says on the tin
	MsgBurnProof = types.MsgBurnProof
)
