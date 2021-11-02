package nft

import (
	"github.com/irisnet/core-sdk-go/types/errors"
)

const CodeSpace = ModuleName

var (
	ErrQueryAddress       = errors.Register(CodeSpace, 1, "query address error")
	ErrBuildAndSend       = errors.Register(CodeSpace, 2, "BuildAndSend error")
	ErrValidateAccAddress = errors.Register(CodeSpace, 3, "Validate account address error")
	ErrQueryNft           = errors.Register(CodeSpace, 4, "QueryNft error")
	ErrDenom              = errors.Register(CodeSpace, 5, "denom  invalid")
	ErrValidateBasic      = errors.Register(CodeSpace, 6, "ValidateBasic fail")
	ErrQueryParams        = errors.Register(CodeSpace, 7, "query params  error")
	ErrGenConn            = errors.Register(CodeSpace, 8, "generate conn error")
)
