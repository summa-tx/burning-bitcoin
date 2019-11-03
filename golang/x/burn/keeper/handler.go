package keeper

import (
	"bytes"
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc"

	"github.com/summa-tx/bitcoin-spv/golang/btcspv"
	"github.com/summa-tx/burning-bitcoin/golang/x/burn/types"
)

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper, ibcKeeper ibc.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgBurnProof:
			return handleMsgBurnProof(ctx, keeper, ibcKeeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to burn proof
func handleMsgBurnProof(ctx sdk.Context, keeper Keeper, ibcKeeper ibc.Keeper, msg types.MsgBurnProof) sdk.Result {
	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic("bad bech32 address in Address")
	}

	value := getBurnInfo(msg.Proof.Vout)
	coin := sdk.NewCoin("burned BTC", sdk.NewInt(int64(value)))

	// TODO: hook in IBC
	ibcKeeper.TransferKeeper.SendTransfer(
		ctx,
		"bank",
		"connectionzero",
		sdk.NewCoins(coin),
		addr,
		addr,
		false)

	keeper.setValidated(ctx, msg.Proof)
	keeper.AppendAddr(ctx, addr)

	return sdk.Result{}
}

func getBurnInfo(vout []byte) uint {
	burnt, _ := hex.DecodeString("759d6677091e973b9e9d99f19c68fbf43e3f05f9")
	numOutputs := uint8(vout[0])
	var o []byte
	for i := uint8(0); i < numOutputs; i++ {
		o, _ = btcspv.ExtractOutputAtIndex(vout, i)
		value, outputType, payload := btcspv.ParseOutput(o)
		if bytes.Equal(payload, burnt) && outputType == btcspv.PKH {
			return value
		}
	}
	return 0
}
