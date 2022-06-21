package slashing

import sdk "github.com/irisnet/core-sdk-go/types"

const (
	TypeMsgUnjailValidator = "unjail_validator"

	ModuleName = "slashing"
)

var (
	_ sdk.Msg = &MsgUnjailValidator{}
)

func (s MsgUnjailValidator) Route() string {
	return ModuleName
}

func (s MsgUnjailValidator) Type() string {
	return TypeMsgUnjailValidator
}

func (s MsgUnjailValidator) ValidateBasic() error {
	if len(s.Id) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "validator id cannot be blank")
	}
	if len(s.Operator) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "operator missing")
	}
	return nil
}

func (s MsgUnjailValidator) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&s)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (s MsgUnjailValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(s.Operator)}
}
