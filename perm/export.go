package perm

import (
	sdk "github.com/irisnet/core-sdk-go/types"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// Client export a group api for Admin module
type Client interface {
	sdk.Module

	AssignRoles(address string, roles []Role, baseTx sdk.BaseTx) (ctypes.ResultTx, error)
	UnassignRoles(address string, roles []Role, baseTx sdk.BaseTx) (ctypes.ResultTx, error)
	BlockAccount(address string, baseTx sdk.BaseTx) (ctypes.ResultTx, error)
	UnblockAccount(address string, baseTx sdk.BaseTx) (ctypes.ResultTx, error)

	QueryRoles(address string) ([]Role, error)
	QueryBlacklist(page, limit int) ([]string, error)
}
