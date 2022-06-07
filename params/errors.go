package params

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

const CodeSpace = ModuleName

var (
	ErrQueryAddress  = sdk.Wrapf(CodeSpace, 1, "query address error")
	ErrBuildAndSend  = sdk.Wrapf(CodeSpace, 2, "BuildAndSend error")
	ErrValidateBasic = sdk.Wrapf(CodeSpace, 3, "ValidateBasic fail")
)
