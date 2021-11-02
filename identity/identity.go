package identity

import (
	"context"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/irisnet/core-sdk-go/codec"
	"github.com/irisnet/core-sdk-go/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/errors"
)

type identityClient struct {
	sdk.BaseClient
	codec.Codec
}

func NewClient(bc sdk.BaseClient, cdc codec.Codec) Client {
	return identityClient{
		BaseClient: bc,
		Codec:      cdc,
	}
}

func (i identityClient) Name() string {
	return ModuleName
}

func (i identityClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (i identityClient) CreateIdentity(request CreateIdentityRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := i.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrQueryAddress, err.Error())
	}

	id, e := sdk.HexBytesFrom(request.ID)
	if e != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrHex, e.Error())
	}

	var pukKeyInfo *PubKeyInfo
	if request.PubkeyInfo != nil {
		if len(request.PubkeyInfo.PubKey) > 0 {
			pubkey, e := sdk.HexBytesFrom(request.PubkeyInfo.PubKey)
			if e != nil {
				return ctypes.ResultTx{}, errors.Wrap(ErrHex, err.Error())
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
		return ctypes.ResultTx{}, errors.Wrap(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (i identityClient) UpdateIdentity(request UpdateIdentityRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := i.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrQueryAddress, err.Error())
	}

	id, e := sdk.HexBytesFrom(request.ID)
	if e != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrHex, e.Error())
	}

	var pukKeyInfo PubKeyInfo
	if request.PubkeyInfo != nil {
		if len(request.PubkeyInfo.PubKey) > 0 {
			pubkey, e := sdk.HexBytesFrom(request.PubkeyInfo.PubKey)
			if e != nil {
				return ctypes.ResultTx{}, errors.Wrap(ErrHex, err.Error())
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
		return ctypes.ResultTx{}, errors.Wrap(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (i identityClient) QueryIdentity(id string) (QueryIdentityResp, error) {
	conn, err := i.GenConn()
	if err != nil {
		return QueryIdentityResp{}, errors.Wrap(ErrGenConn, err.Error())
	}

	identityId, err := sdk.HexBytesFrom(id)
	if err != nil {
		return QueryIdentityResp{}, errors.Wrap(ErrHex, err.Error())
	}

	resp, err := NewQueryClient(conn).Identity(
		context.Background(),
		&QueryIdentityRequest{Id: identityId.String()},
	)
	if err != nil {
		return QueryIdentityResp{}, errors.Wrap(ErrQueryIdentity, err.Error())
	}

	return resp.Identity.Convert().(QueryIdentityResp), nil
}
