package perm

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

const CodeSpace = ModuleName

var (
	ErrQueryAddress  = sdk.Wrapf(CodeSpace, 1, "query address error")
	ErrBuildAndSend  = sdk.Wrapf(CodeSpace, 2, "BuildAndSend error")
	ErrGenConn       = sdk.Wrapf(CodeSpace, 3, "generate conn error")
	ErrBench32       = sdk.Wrapf(CodeSpace, 4, "err bench32")
	ErrQueryPerm     = sdk.Wrapf(CodeSpace, 5, "query perm fail")
	ErrValidateBasic = sdk.Wrapf(CodeSpace, 6, "ValidateBasic fail")
)
