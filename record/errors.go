package record

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

const CodeSpace = ModuleName

var (
	ErrQueryAddress  = sdk.Wrapf(CodeSpace, 1, "query address error")
	ErrBuildAndSend  = sdk.Wrapf(CodeSpace, 2, "BuildAndSend error")
	ErrValidateBasic = sdk.Wrapf(CodeSpace, 3, "ValidateBasic fail")
	ErrHex           = sdk.Wrapf(CodeSpace, 4, "hex fail")
	ErrUnmarshal     = sdk.Wrapf(CodeSpace, 5, "unmarshal fail")
	ErrGetEvents     = sdk.Wrapf(CodeSpace, 6, "get events fail")
	ErrQueryStore    = sdk.Wrapf(CodeSpace, 7, "QueryStore error")
)
