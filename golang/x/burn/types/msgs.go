package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RouterKey is a name for the router
const RouterKey = ModuleName // this was defined in your key.go file

// MsgDoThing defines a DoThing message
type MsgDoThing struct {
}

// NewMsgDoThing is a constructor function for MsgDoThing
func NewMsgDoThing(owner sdk.AccAddress) MsgDoThing {
	return MsgDoThing{}
}

// // MUST IMPLEMENT

// GetSigners defines whose signature is required
func (msg MsgDoThing) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

// Type should return the action
func (msg MsgDoThing) Type() string { return "do_thing" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDoThing) ValidateBasic() sdk.Error {
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDoThing) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// Route should return the name of the module
func (msg MsgDoThing) Route() string { return RouterKey }
