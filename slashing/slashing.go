package slashing

import (
	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type slashingClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) Client {
	return slashingClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (s slashingClient) Name() string {
	return ModuleName
}

func (s slashingClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (s slashingClient) UnjailValidator(id string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgUnjailValidator{
		Id:       id,
		Operator: sender.String(),
	}
	res, err := s.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return res, nil
}
