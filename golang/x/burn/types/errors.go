package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// DefaultCodespace is the default code space
	DefaultCodespace sdk.CodespaceType = ModuleName

	// SomeProblem is a problem!!!
	SomeProblem sdk.CodeType = 999
)

// ErrSomeProblem throws an error
func ErrSomeProblem(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, SomeProblem, "UH OH")
}
