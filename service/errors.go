package service

import (
	"github.com/irisnet/core-sdk-go/types/errors"
)

const CodeSpace = ModuleName

var (
	ErrQueryAddress       = errors.Register(CodeSpace, 1, "query address error")
	ErrBuildAndSend       = errors.Register(CodeSpace, 2, "BuildAndSend error")
	ErrValidateBasic      = errors.Register(CodeSpace, 3, "ValidateBasic fail")
	ErrHex                = errors.Register(CodeSpace, 4, "hex fail")
	ErrGenConn            = errors.Register(CodeSpace, 5, "generate conn error")
	ErrRequest            = errors.Register(CodeSpace, 6, "request fail")
	ErrContext            = errors.Register(CodeSpace, 7, "Context fail")
	ErrQueryTx            = errors.Register(CodeSpace, 8, "Query tx fail")
	ErrTxDecoder          = errors.Register(CodeSpace, 9, "TxDecoder fail")
	ErrQueryBlock         = errors.Register(CodeSpace, 10, "query block fail")
	ErrUnmarshal          = errors.Register(CodeSpace, 11, "unmarshal fail")
	ErrMsg                = errors.Register(CodeSpace, 12, "Msg invalid")
	ErrValidateAccAddress = errors.Register(CodeSpace, 13, "Validate account address error")
)
