package identity

import (
	"github.com/irisnet/core-sdk-go/types/errors"
)

const CodeSpace = ModuleName

var (
	ErrQueryAddress  = errors.Register(CodeSpace, 1, "query address error")
	ErrBuildAndSend  = errors.Register(CodeSpace, 2, "BuildAndSend error")
	ErrQueryIdentity = errors.Register(CodeSpace, 3, "QueryIdentity error")
	ErrGenConn       = errors.Register(CodeSpace, 4, "generate conn error")
	ErrHex           = errors.Register(CodeSpace, 5, "hex fail")
	ErrValidateBasic = errors.Register(CodeSpace, 6, "ValidateBasic fail")
	ErrUnmarshal     = errors.Register(CodeSpace, 7, "unmarshal fail")
)
