package identity

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

const CodeSpace = ModuleName

var (
	ErrQueryAddress  = sdk.Wrapf("query address error")
	ErrBuildAndSend  = sdk.Wrapf("BuildAndSend error")
	ErrQueryIdentity = sdk.Wrapf("QueryIdentity error")
	ErrGenConn       = sdk.Wrapf("generate conn error")
	ErrHex           = sdk.Wrapf("hex fail")
	ErrValidateBasic = sdk.Wrapf("ValidateBasic fail")
	ErrUnmarshal     = sdk.Wrapf("unmarshal fail")
)
