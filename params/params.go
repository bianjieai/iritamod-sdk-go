package params

import (
	"fmt"

	"github.com/irisnet/core-sdk-go/types/errors"

	"github.com/irisnet/core-sdk-go/codec"
	"github.com/irisnet/core-sdk-go/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

type paramsClient struct {
	sdk.BaseClient
	codec.Codec
}

func NewClient(bc sdk.BaseClient, cdc codec.Codec) Client {
	return paramsClient{
		BaseClient: bc,
		Codec:      cdc,
	}
}

func (p paramsClient) Name() string {
	return ModuleName
}

func (p paramsClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (p paramsClient) UpdateParams(requests []UpdateParamRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := p.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrapf(ErrQueryAddress, fmt.Sprintf("%s not found", baseTx.From))
	}

	var changes []ParamChange
	for _, req := range requests {
		changes = append(changes, ParamChange{
			Subspace: req.Module,
			Key:      req.Key,
			Value:    req.Value,
		})
	}

	msg := &MsgUpdateParams{
		Changes:  changes,
		Operator: sender.String(),
	}
	send, err := p.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrBuildAndSend, err.Error())
	}
	return send, nil
}
