package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	restName = "burn"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/burnProof", storeName), burnProofHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/validated/{%s}", storeName, restName), validatedHandler(cliCtx, storeName)).Methods("GET")
}
