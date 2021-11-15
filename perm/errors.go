package perm

import (
	"github.com/irisnet/core-sdk-go/types/errors"
)

const CodeSpace = ModuleName

var (
	ErrQueryAddress  = errors.Register(CodeSpace, 1, "query address error")
	ErrBuildAndSend  = errors.Register(CodeSpace, 2, "BuildAndSend error")
	ErrGenConn       = errors.Register(CodeSpace, 3, "generate conn error")
	ErrBench32       = errors.Register(CodeSpace, 4, "err bench32")
	ErrQueryPerm     = errors.Register(CodeSpace, 5, "query perm fail")
	ErrValidateBasic = errors.Register(CodeSpace, 6, "ValidateBasic fail")
)
