package params

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/errors"
)

const (
	ModuleName = "params"
)

var (
	_ sdk.Msg = &MsgUpdateParams{}
)

func (m MsgUpdateParams) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return errors.Wrap(ErrValidateBasic, "operator missing")
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
		return errors.Wrap(ErrValidateBasic, "no change params")
	}

	for _, pc := range changes {
		if len(pc.Subspace) == 0 {
			return errors.Wrap(ErrValidateBasic, "empty subspace")
		}
		if len(pc.Key) == 0 {
			return errors.Wrap(ErrValidateBasic, "empty params key")
		}
		if len(pc.Value) == 0 {
			return errors.Wrap(ErrValidateBasic, "empty params value")
		}
	}

	return nil
}
