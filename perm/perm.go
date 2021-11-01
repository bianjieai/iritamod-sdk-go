package perm

import (
	"context"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/irisnet/core-sdk-go/codec"
	"github.com/irisnet/core-sdk-go/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type permClient struct {
	sdk.BaseClient
	codec.Codec
}

func NewClient(bc sdk.BaseClient, cdc codec.Codec) Client {
	return permClient{
		BaseClient: bc,
		Codec:      cdc,
	}
}

func (a permClient) Name() string {
	return ModuleName
}

func (a permClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (a permClient) AssignRoles(address string, roles []Role, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := a.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	acc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	msg := &MsgAssignRoles{
		Address:  acc.String(),
		Roles:    roles,
		Operator: sender.String(),
	}
	return a.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (a permClient) UnassignRoles(address string, roles []Role, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := a.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	acc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	msg := &MsgUnassignRoles{
		Address:  acc.String(),
		Roles:    roles,
		Operator: sender.String(),
	}
	return a.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (a permClient) BlockAccount(address string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := a.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	acc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	msg := &MsgBlockAccount{
		Address:  acc.String(),
		Operator: sender.String(),
	}
	return a.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (a permClient) UnblockAccount(address string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := a.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	acc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	msg := &MsgUnblockAccount{
		Address:  acc.String(),
		Operator: sender.String(),
	}
	return a.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (a permClient) QueryRoles(address string) ([]Role, error) {
	conn, err := a.GenConn()

	if err != nil {
		return nil, err
	}

	acc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return nil, err
	}

	resp, err := NewQueryClient(conn).Roles(
		context.Background(),
		&QueryRolesRequest{Address: acc.String()},
	)
	if err != nil {
		return nil, err
	}

	return resp.Roles, nil
}

func (a permClient) QueryBlacklist(page, limit int) ([]string, error) {
	conn, err := a.GenConn()

	if err != nil {
		return nil, err
	}

	resp, err := NewQueryClient(conn).Blacklist(
		context.Background(),
		&QueryBlacklistRequest{},
	)
	if err != nil {
		return nil, err
	}

	return resp.Addresses, nil
}
