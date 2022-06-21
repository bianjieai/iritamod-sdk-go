package perm

import (
	"context"

	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type permClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) Client {
	return permClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (a permClient) Name() string {
	return ModuleName
}

func (a permClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (a permClient) AssignRoles(address string, roles []Role, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := a.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	acc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBench32, err.Error())
	}

	msg := &MsgAssignRoles{
		Address:  acc.String(),
		Roles:    roles,
		Operator: sender.String(),
	}
	send, err := a.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (a permClient) UnassignRoles(address string, roles []Role, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := a.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	acc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBench32, err.Error())
	}

	msg := &MsgUnassignRoles{
		Address:  acc.String(),
		Roles:    roles,
		Operator: sender.String(),
	}

	send, err := a.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (a permClient) BlockAccount(address string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := a.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	acc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBench32, err.Error())
	}

	msg := &MsgBlockAccount{
		Address:  acc.String(),
		Operator: sender.String(),
	}

	send, err := a.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (a permClient) UnblockAccount(address string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := a.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	acc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBench32, err.Error())
	}

	msg := &MsgUnblockAccount{
		Address:  acc.String(),
		Operator: sender.String(),
	}

	send, err := a.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (a permClient) BlockContract(address string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := a.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	acc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBench32, err.Error())
	}

	msg := &MsgBlockAccount{
		Address:  acc.String(),
		Operator: sender.String(),
	}

	send, err := a.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (a permClient) UnblockContract(address string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := a.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	acc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBench32, err.Error())
	}

	msg := &MsgUnblockAccount{
		Address:  acc.String(),
		Operator: sender.String(),
	}

	send, err := a.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (a permClient) QueryRoles(address string) ([]Role, error) {
	conn, err := a.GenConn()

	if err != nil {
		return nil, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}

	acc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrBench32, err.Error())
	}

	resp, err := NewQueryClient(conn).Roles(
		context.Background(),
		&QueryRolesRequest{Address: acc.String()},
	)
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrQueryPerm, err.Error())
	}

	return resp.Roles, nil
}

func (a permClient) QueryAccountBlockList() ([]string, error) {
	conn, err := a.GenConn()

	if err != nil {
		return nil, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}

	resp, err := NewQueryClient(conn).AccountBlockList(
		context.Background(),
		&QueryBlockListRequest{},
	)
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrQueryPerm, err.Error())
	}

	return resp.Addresses, nil
}

func (a permClient) QueryContractDenyList() ([]string, error) {
	conn, err := a.GenConn()

	if err != nil {
		return nil, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}

	resp, err := NewQueryClient(conn).ContractDenyList(
		context.Background(),
		&QueryContractDenyList{},
	)
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrQueryPerm, err.Error())
	}

	return resp.Addresses, nil
}
