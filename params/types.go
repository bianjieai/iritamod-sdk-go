package params

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

const (
	ModuleName = "params"
)

var (
	_ sdk.Msg = &MsgUpdateParams{}
)

func (m MsgUpdateParams) Route() string {
	return ModuleName
}

func (m MsgUpdateParams) Type() string {
	return "update_params"
}

func (m MsgUpdateParams) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgUpdateParams) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "operator missing")
	}
	return validateChanges(m.Changes)
}

func (m MsgUpdateParams) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

// ValidateChanges performs basic validation checks over a set of ParamChange. It
// returns an error if any ParamChange is invalid.
func validateChanges(changes []ParamChange) error {
	if len(changes) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "no change params")
	}

	for _, pc := range changes {
		if len(pc.Subspace) == 0 {
			return sdk.WrapWithMessage(ErrValidateBasic, "empty subspace")
		}
		if len(pc.Key) == 0 {
			return sdk.WrapWithMessage(ErrValidateBasic, "empty params key")
		}
		if len(pc.Value) == 0 {
			return sdk.WrapWithMessage(ErrValidateBasic, "empty params value")
		}
	}

	return nil
}
