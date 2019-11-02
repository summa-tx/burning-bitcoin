package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/summa-tx/bitcoin-spv/golang/btcspv"
)

// RouterKey is a name for the router
const RouterKey = ModuleName // this was defined in your key.go file

// MsgBurnProof represents a proof of burn
type MsgBurnProof struct {
	Proof       btcspv.SPVProof        `json:"proof"`
	HeaderChain []btcspv.BitcoinHeader `json:"headers"`
	Address     string                 `json:"signer"`
}

// NewMsgBurnProof is a constructor function for MsgBurnProof
func NewMsgBurnProof(proof btcspv.SPVProof, chain []btcspv.BitcoinHeader, address string, owner sdk.AccAddress) MsgBurnProof {
	return MsgBurnProof{
		proof,
		chain,
		address,
	}
}

// GetSigners defines whose signature is required
func (msg MsgBurnProof) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		// Nobody is authed
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{addr}
}

// Type should return the action
func (msg MsgBurnProof) Type() string { return "burnProof" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBurnProof) ValidateBasic() sdk.Error {
	validHeader, err := msg.Proof.ConfirmingHeader.Validate()
	if !validHeader {
		return ErrBadHeader(err.Error(), DefaultCodespace)
	}

	validProof, err := msg.Proof.Validate()
	if !validProof {
		return ErrBadProof(err.Error(), DefaultCodespace)
	}

	validChain, err := validateChain(msg.Proof.ConfirmingHeader, msg.HeaderChain)
	if validChain.IsZero() {
		return ErrBadChain(err.Error(), DefaultCodespace)
	}

	return nil
}

// validate the header chain
func validateChain(first btcspv.BitcoinHeader, chain []btcspv.BitcoinHeader) (sdk.Uint, error) {
	buf := make([]byte, 80*(1+len(chain)))
	buf = append(buf, first.Raw[:]...)
	for i := 0; i < len(chain); i++ {
		buf = append(buf, chain[i].Raw[:]...)
	}

	return btcspv.ValidateHeaderChain(buf)
}

// GetSignBytes encodes the message for signing
func (msg MsgBurnProof) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// Route should return the name of the module
func (msg MsgBurnProof) Route() string { return RouterKey }
