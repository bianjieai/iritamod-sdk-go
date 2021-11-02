package token

import (
	"github.com/irisnet/core-sdk-go/types/errors"
)

const CodeSpace = ModuleName

var (
	ErrQueryAddress       = errors.Register(CodeSpace, 1, "query address error")
	ErrBuildAndSend       = errors.Register(CodeSpace, 2, "BuildAndSend error")
	ErrValidateBasic      = errors.Register(CodeSpace, 3, "ValidateBasic fail")
	ErrQueryToken         = errors.Register(CodeSpace, 4, "QueryToken error")
	ErrGenConn            = errors.Register(CodeSpace, 5, "generate conn error")
	ErrValidateAccAddress = errors.Register(CodeSpace, 6, "Validate account address error")
	ErrUnmarshal          = errors.Register(CodeSpace, 7, "unmarshal fail")
)
