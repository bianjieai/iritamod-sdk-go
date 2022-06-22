package node

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

const CodeSpace = ModuleName

var (
	ErrQueryAddress  = sdk.Wrapf(CodeSpace, 1, "query address error")
	ErrBuildAndSend  = sdk.Wrapf(CodeSpace, 2, "BuildAndSend error")
	ErrValidateBasic = sdk.Wrapf(CodeSpace, 3, "ValidateBasic fail")
	ErrQueryNode     = sdk.Wrapf(CodeSpace, 4, "QueryNode error")
	ErrGenConn       = sdk.Wrapf(CodeSpace, 5, "generate conn error")
	ErrHex           = sdk.Wrapf(CodeSpace, 6, "hex fail")
)
