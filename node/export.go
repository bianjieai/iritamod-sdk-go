package node

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/query"
)

var (
	_ Client = nodeClient{}
)

// expose Record module api for user
type Client interface {
	sdk.Module

	CreateValidator(request CreateValidatorRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	UpdateValidator(request UpdateValidatorRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	RemoveValidator(id string, baseTx sdk.BaseTx) (sdk.ResultTx, error)

	GrantNode(request GrantNodeRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	RevokeNode(nodeId string, baseTx sdk.BaseTx) (sdk.ResultTx, error)

	QueryValidators(pageReq *query.PageRequest) ([]QueryValidatorResp, error)
	QueryValidator(id string) (QueryValidatorResp, error)
	QueryNodes(pageReq *query.PageRequest) ([]QueryNodeResp, error)
	QueryNode(id string) (QueryNodeResp, error)
	QueryParams() (QueryParamsResp, error)
}

type CreateValidatorRequest struct {
	Name        string `json:"name"`
	Certificate string `json:"certificate"`
	Power       int64  `json:"power"`
	Details     string `json:"details"`
}

type UpdateValidatorRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Certificate string `json:"certificate"`
	Power       int64  `json:"power"`
	Details     string `json:"details"`
}

type GrantNodeRequest struct {
	Name        string `json:"name"`
	Certificate string `json:"certificate"`
	Details     string `json:"details"`
}

type QueryValidatorResp struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Pubkey      string `json:"pubkey"`
	Certificate string `json:"certificate"`
	Power       int64  `json:"power"`
	Details     string `json:"details"`
	Jailed      bool   `json:"jailed"`
	Operator    string `json:"operator"`
}

type QueryNodeResp struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Certificate string `json:"certificate"`
}

// token params
type QueryParamsResp struct {
	HistoricalEntries uint32 `json:"historical_entries"`
}
