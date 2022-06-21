package identity

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

const CodeSpace = ModuleName

var (
	ErrQueryAddress  = sdk.Wrapf(CodeSpace, 1, "query address error")
	ErrBuildAndSend  = sdk.Wrapf(CodeSpace, 2, "BuildAndSend error")
	ErrQueryIdentity = sdk.Wrapf(CodeSpace, 3, "QueryIdentity error")
	ErrGenConn       = sdk.Wrapf(CodeSpace, 4, "generate conn error")
	ErrHex           = sdk.Wrapf(CodeSpace, 5, "hex fail")
	ErrValidateBasic = sdk.Wrapf(CodeSpace, 6, "ValidateBasic fail")
	ErrUnmarshal     = sdk.Wrapf(CodeSpace, 7, "unmarshal fail")
)
