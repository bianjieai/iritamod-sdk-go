package identity

import (
	"context"

	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type identityClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) Client {
	return identityClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (i identityClient) Name() string {
	return ModuleName
}

func (i identityClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (i identityClient) CreateIdentity(request CreateIdentityRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := i.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	id, e := sdk.HexBytesFrom(request.ID)
	if e != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrHex, e.Error())
	}

	var pukKeyInfo *PubKeyInfo
	if request.PubkeyInfo != nil {
		if len(request.PubkeyInfo.PubKey) > 0 {
			pubkey, e := sdk.HexBytesFrom(request.PubkeyInfo.PubKey)
			if e != nil {
				return sdk.ResultTx{}, sdk.WrapWithMessage(ErrHex, err.Error())
			}
			pukKeyInfo = &PubKeyInfo{
				PubKey:    pubkey.String(),
				Algorithm: request.PubkeyInfo.PubKeyAlgo,
			}
		}
	}

	credentials := ""
	if request.Credentials != nil {
		credentials = *request.Credentials
	}
	msg := &MsgCreateIdentity{
		Id:          id.String(),
		PubKey:      pukKeyInfo,
		Certificate: request.Certificate,
		Credentials: credentials,
		Owner:       sender.String(),
	}
	send, err := i.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (i identityClient) UpdateIdentity(request UpdateIdentityRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := i.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	id, e := sdk.HexBytesFrom(request.ID)
	if e != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrHex, e.Error())
	}

	var pukKeyInfo PubKeyInfo
	if request.PubkeyInfo != nil {
		if len(request.PubkeyInfo.PubKey) > 0 {
			pubkey, e := sdk.HexBytesFrom(request.PubkeyInfo.PubKey)
			if e != nil {
				return sdk.ResultTx{}, sdk.WrapWithMessage(ErrHex, err.Error())
			}
			pukKeyInfo.PubKey = pubkey.String()
			pukKeyInfo.Algorithm = request.PubkeyInfo.PubKeyAlgo
		}
	}

	credentials := DoNotModifyDesc
	if request.Credentials != nil {
		credentials = *request.Credentials
	}

	msg := &MsgUpdateIdentity{
		Id:          id.String(),
		PubKey:      &pukKeyInfo,
		Certificate: request.Certificate,
		Credentials: credentials,
		Owner:       sender.String(),
	}
	send, err := i.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (i identityClient) QueryIdentity(id string) (QueryIdentityResp, error) {
	conn, err := i.GenConn()
	if err != nil {
		return QueryIdentityResp{}, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}

	identityId, err := sdk.HexBytesFrom(id)
	if err != nil {
		return QueryIdentityResp{}, sdk.WrapWithMessage(ErrHex, err.Error())
	}

	resp, err := NewQueryClient(conn).Identity(
		context.Background(),
		&QueryIdentityRequest{Id: identityId.String()},
	)
	if err != nil {
		return QueryIdentityResp{}, sdk.WrapWithMessage(ErrQueryIdentity, err.Error())
	}

	return resp.Identity.Convert().(QueryIdentityResp), nil
}
