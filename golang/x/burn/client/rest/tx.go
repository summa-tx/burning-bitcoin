package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/summa-tx/burning-bitcoin/golang/x/burn/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/summa-tx/bitcoin-spv/golang/btcspv"
)

type burnProofReq struct {
	BaseReq     rest.BaseReq           `json:"base_req"`
	Proof       btcspv.SPVProof        `json:"proof"`
	HeaderChain []btcspv.BitcoinHeader `json:"headers"`
	Address     string                 `json:"signer"`
}

func burnProofHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req burnProofReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgBurnProof(req.Proof, req.HeaderChain, req.Address, sdk.AccAddress{})
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
