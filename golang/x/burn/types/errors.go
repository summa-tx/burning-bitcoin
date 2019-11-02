package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// DefaultCodespace is the default code space
	DefaultCodespace sdk.CodespaceType = ModuleName

	// BadHeader is a bad spv proof header
	BadHeader = 333

	// BadChain is a bad header chain
	BadChain = 222

	// BadProof is a bad SPV Proof
	BadProof = 111

	// SomeProblem is a problem!!!
	SomeProblem sdk.CodeType = 999
)

// ErrSomeProblem throws an error
func ErrSomeProblem(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, SomeProblem, "UH OH")
}

// ErrBadHeader throws an error
func ErrBadHeader(reason string, codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, BadHeader, reason)
}

// ErrBadProof throws an error
func ErrBadProof(reason string, codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, BadProof, reason)
}

// ErrBadChain throws an error
func ErrBadChain(reason string, codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, BadChain, reason)
}
