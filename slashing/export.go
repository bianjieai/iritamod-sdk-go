package slashing

import sdk "github.com/irisnet/core-sdk-go/types"

type Client interface {
	sdk.Module

	UnjailValidator(id string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
}
