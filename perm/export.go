package perm

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

// Client export a group api for Admin module
type Client interface {
	sdk.Module

	AssignRoles(address string, roles []Role, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	UnassignRoles(address string, roles []Role, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	BlockAccount(address string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	UnblockAccount(address string, baseTx sdk.BaseTx) (sdk.ResultTx, error)

	QueryRoles(address string) ([]Role, error)
	QueryBlacklist(page, limit int) ([]string, error)
}
