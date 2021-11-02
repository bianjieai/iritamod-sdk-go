package wasm

import (
	"github.com/irisnet/core-sdk-go/types/errors"
)

const CodeSpace = ModuleName

var (
	ErrQueryAddress       = errors.Register(CodeSpace, 1, "query address error")
	ErrBuildAndSend       = errors.Register(CodeSpace, 2, "BuildAndSend error")
	ErrValidateBasic      = errors.Register(CodeSpace, 3, "ValidateBasic fail")
	ErrQueryWasm          = errors.Register(CodeSpace, 4, "query wasm fail")
	ErrGenConn            = errors.Register(CodeSpace, 5, "generate conn error")
	ErrValidateAccAddress = errors.Register(CodeSpace, 6, "Validate account address error")
	ErrAbiBuild           = errors.Register(CodeSpace, 7, "abi build fail")
	ErrReadFile           = errors.Register(CodeSpace, 8, "read file fail")
	ErrGetEvents          = errors.Register(CodeSpace, 9, "get events fail")
)
