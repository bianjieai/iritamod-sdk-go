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

	// expect request.ID is hex-able.
	_, hexErr := sdk.HexBytesFrom(request.Id)
	if hexErr != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrHex, hexErr.Error())
	}

	// expect pubkey is hex-able
	if request.PubKeyInfo != nil && len(request.PubKeyInfo.PubKey) > 0 {
		_, hexErr := sdk.HexBytesFrom(request.PubKeyInfo.PubKey)
		if hexErr != nil {
			return sdk.ResultTx{}, sdk.WrapWithMessage(ErrHex, err.Error())
		}
	}

	msg := &MsgCreateIdentity{
		Id:          request.Id,
		PubKey:      request.PubKeyInfo,
		Certificate: request.Certificate,
		Credentials: *request.Credentials,
		Owner:       sender.String(),
		Data:        request.Data,
	}

	res, err := i.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return res, nil
}

func (i identityClient) UpdateIdentity(request UpdateIdentityRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := i.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	id, e := sdk.HexBytesFrom(request.Id)
	if e != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrHex, e.Error())
	}

	if request.PubKeyInfo != nil && len(request.PubKeyInfo.PubKey) > 0 {
		_, hexErr := sdk.HexBytesFrom(request.PubKeyInfo.PubKey)
		if hexErr != nil {
			return sdk.ResultTx{}, sdk.WrapWithMessage(ErrHex, err.Error())
		}
	}

	credentials := DoNotModifyDesc
	if request.Credentials != nil {
		credentials = *request.Credentials
	}
	data := DoNotModifyDesc
	if request.Data != nil {
		data = *request.Data
	}

	msg := &MsgUpdateIdentity{
		Id:          id.String(),
		PubKey:      request.PubKeyInfo,
		Certificate: request.Certificate,
		Credentials: credentials,
		Owner:       sender.String(),
		Data:        data,
	}

	res, err := i.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return res, nil
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
