package record

import (
	"github.com/irisnet/core-sdk-go/types/errors"
)

const CodeSpace = ModuleName

var (
	ErrQueryAddress  = errors.Register(CodeSpace, 1, "query address error")
	ErrBuildAndSend  = errors.Register(CodeSpace, 2, "BuildAndSend error")
	ErrValidateBasic = errors.Register(CodeSpace, 3, "ValidateBasic fail")
	ErrHex           = errors.Register(CodeSpace, 4, "hex fail")
	ErrUnmarshal     = errors.Register(CodeSpace, 5, "unmarshal fail")
	ErrGetEvents     = errors.Register(CodeSpace, 6, "get events fail")
	ErrQueryStore    = errors.Register(CodeSpace, 7, "QueryStore error")
)
