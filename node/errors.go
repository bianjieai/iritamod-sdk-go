package node

import (
	"github.com/irisnet/core-sdk-go/types/errors"
)

const CodeSpace = ModuleName

var (
	ErrQueryAddress  = errors.Register(CodeSpace, 1, "query address error")
	ErrBuildAndSend  = errors.Register(CodeSpace, 2, "BuildAndSend error")
	ErrValidateBasic = errors.Register(CodeSpace, 3, "ValidateBasic fail")
	ErrQueryNode     = errors.Register(CodeSpace, 4, "QueryNode error")
	ErrGenConn       = errors.Register(CodeSpace, 5, "generate conn error")
	ErrHex           = errors.Register(CodeSpace, 6, "hex fail")
)
