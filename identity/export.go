package identity

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

type Client interface {
	sdk.Module

	CreateIdentity(request CreateIdentityRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	UpdateIdentity(request UpdateIdentityRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error)

	QueryIdentity(id string) (QueryIdentityResp, error)
}

type CreateIdentityRequest struct {
	Id          string      `json:"id"`
	PubKeyInfo  *PubKeyInfo `json:"pubkey_info"`
	Certificate string      `json:"certificate"`
	Credentials *string     `json:"credentials"`
	Data        string      `json:"data"`
}

type UpdateIdentityRequest struct {
	Id          string      `json:"id"`
	PubKeyInfo  *PubKeyInfo `json:"pubkey_info"`
	Certificate string      `json:"certificate"`
	Credentials *string     `json:"credentials"`
	Data        *string     `json:"data"`
}

type QueryIdentityResp struct {
	Id           string       `json:"id"`
	PubKeyInfos  []PubKeyInfo `json:"pubkey_infos"`
	Certificates []string     `json:"certificates"`
	Credentials  string       `json:"credentials"`
	Owner        string       `json:"owner"`
	Data         string       `json:"data"`
}
