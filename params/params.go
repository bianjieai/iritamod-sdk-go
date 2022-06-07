package params

import (
	"fmt"

	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type paramsClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) Client {
	return paramsClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (p paramsClient) Name() string {
	return ModuleName
}

func (p paramsClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (p paramsClient) UpdateParams(requests []UpdateParamRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := p.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, fmt.Sprintf("%s not found", baseTx.From))
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
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return send, nil
}
