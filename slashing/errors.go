package slashing

import sdk "github.com/irisnet/core-sdk-go/types"

const CodeSpace = ModuleName

var (
	ErrQueryAddress  = sdk.Wrapf(CodeSpace, 1, "query address error")
	ErrBuildAndSend  = sdk.Wrapf(CodeSpace, 2, "BuildAndSend error")
	ErrHex           = sdk.Wrapf(CodeSpace, 3, "hex fail")
	ErrValidateBasic = sdk.Wrapf(CodeSpace, 4, "ValidateBasic fail")
)
