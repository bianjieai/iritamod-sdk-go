package token

import sdk "github.com/irisnet/core-sdk-go/types"

const CodeSpace = ModuleName

var (
	ErrQueryAddress       = sdk.Wrapf(CodeSpace, 1, "query address error")
	ErrBuildAndSend       = sdk.Wrapf(CodeSpace, 2, "BuildAndSend error")
	ErrValidateBasic      = sdk.Wrapf(CodeSpace, 3, "ValidateBasic fail")
	ErrQueryToken         = sdk.Wrapf(CodeSpace, 4, "QueryToken error")
	ErrGenConn            = sdk.Wrapf(CodeSpace, 5, "generate conn error")
	ErrValidateAccAddress = sdk.Wrapf(CodeSpace, 6, "Validate account address error")
	ErrUnmarshal          = sdk.Wrapf(CodeSpace, 7, "unmarshal fail")
)
