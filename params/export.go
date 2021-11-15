package params

import (
	sdk "github.com/irisnet/core-sdk-go/types"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

var (
	_ Client = paramsClient{}
)

// expose params module api for user
type Client interface {
	sdk.Module

	UpdateParams(request []UpdateParamRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error)
}

type UpdateParamRequest struct {
	Module string `json:"module"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}
