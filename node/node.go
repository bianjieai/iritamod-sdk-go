package node

import (
	"context"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/irisnet/core-sdk-go/codec"
	"github.com/irisnet/core-sdk-go/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/errors"
	query "github.com/irisnet/core-sdk-go/types/query"
)

type nodeClient struct {
	sdk.BaseClient
	codec.Codec
}

func NewClient(bc sdk.BaseClient, cdc codec.Codec) Client {
	return nodeClient{
		BaseClient: bc,
		Codec:      cdc,
	}
}

func (n nodeClient) Name() string {
	return ModuleName
}

func (n nodeClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (n nodeClient) CreateValidator(request CreateValidatorRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	creator, err := n.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrQueryAddress, err.Error())
	}

	msg := &MsgCreateValidator{
		Name:        request.Name,
		Certificate: request.Certificate,
		Description: request.Details,
		Power:       request.Power,
		Operator:    creator.String(),
	}
	send, err := n.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (n nodeClient) UpdateValidator(request UpdateValidatorRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	creator, err := n.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrQueryAddress, err.Error())
	}

	vID, er := sdk.HexBytesFrom(request.ID)
	if er != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrHex, err.Error())
	}

	msg := &MsgUpdateValidator{
		Id:          vID.String(),
		Name:        request.Name,
		Certificate: request.Certificate,
		Description: request.Details,
		Power:       request.Power,
		Operator:    creator.String(),
	}

	send, err := n.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (n nodeClient) RemoveValidator(id string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	creator, err := n.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrQueryAddress, err.Error())
	}

	vID, er := sdk.HexBytesFrom(id)
	if er != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrHex, err.Error())
	}
	msg := &MsgRemoveValidator{
		Id:       vID.String(),
		Operator: creator.String(),
	}

	send, err := n.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (n nodeClient) GrantNode(request GrantNodeRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	creator, err := n.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrQueryAddress, err.Error())
	}

	msg := &MsgGrantNode{
		Name:        request.Name,
		Certificate: request.Certificate,
		Operator:    creator.String(),
	}

	send, err := n.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (n nodeClient) RevokeNode(nodeId string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	creator, err := n.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrQueryAddress, err.Error())
	}

	vID, er := sdk.HexBytesFrom(nodeId)
	if er != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrHex, er.Error())
	}

	msg := &MsgRevokeNode{
		Id:       vID.String(),
		Operator: creator.String(),
	}

	send, err := n.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (n nodeClient) QueryValidators(pageReq *query.PageRequest) ([]QueryValidatorResp, error) {
	conn, err := n.GenConn()

	if err != nil {
		return nil, errors.Wrap(ErrGenConn, err.Error())
	}

	resp, err := NewQueryClient(conn).Validators(
		context.Background(),
		&QueryValidatorsRequest{
			Pagination: pageReq,
		},
	)
	if err != nil {
		return nil, errors.Wrap(ErrQueryNode, err.Error())
	}

	return validators(resp.Validators).Convert().([]QueryValidatorResp), nil
}

func (n nodeClient) QueryValidator(id string) (QueryValidatorResp, error) {
	conn, err := n.GenConn()

	if err != nil {
		return QueryValidatorResp{}, errors.Wrap(ErrGenConn, err.Error())
	}

	vID, err := sdk.HexBytesFrom(id)
	if err != nil {
		return QueryValidatorResp{}, errors.Wrap(ErrHex, err.Error())
	}

	resp, err := NewQueryClient(conn).Validator(
		context.Background(),
		&QueryValidatorRequest{
			Id: vID.String(),
		},
	)
	if err != nil {
		return QueryValidatorResp{}, errors.Wrap(ErrQueryNode, err.Error())
	}

	return resp.Validator.Convert().(QueryValidatorResp), nil
}

func (n nodeClient) QueryNodes(pageReq *query.PageRequest) ([]QueryNodeResp, error) {
	conn, err := n.GenConn()

	if err != nil {
		return nil, errors.Wrap(ErrGenConn, err.Error())
	}

	resp, err := NewQueryClient(conn).Nodes(
		context.Background(),
		&QueryNodesRequest{
			Pagination: pageReq,
		},
	)
	if err != nil {
		return nil, errors.Wrap(ErrQueryNode, err.Error())
	}

	return nodes(resp.Nodes).Convert().([]QueryNodeResp), nil
}

func (n nodeClient) QueryNode(id string) (QueryNodeResp, error) {
	conn, err := n.GenConn()

	if err != nil {
		return QueryNodeResp{}, errors.Wrap(ErrGenConn, err.Error())
	}

	vID, err := sdk.HexBytesFrom(id)
	if err != nil {
		return QueryNodeResp{}, errors.Wrap(ErrHex, err.Error())
	}

	resp, err := NewQueryClient(conn).Node(
		context.Background(),
		&QueryNodeRequest{
			Id: vID.String(),
		},
	)
	if err != nil {
		return QueryNodeResp{}, errors.Wrap(ErrQueryNode, err.Error())
	}

	return resp.Node.Convert().(QueryNodeResp), nil
}

func (n nodeClient) QueryParams() (QueryParamsResp, error) {
	conn, err := n.GenConn()

	if err != nil {
		return QueryParamsResp{}, errors.Wrap(ErrGenConn, err.Error())
	}

	resp, err := NewQueryClient(conn).Params(
		context.Background(),
		&QueryParamsRequest{},
	)
	if err != nil {
		return QueryParamsResp{}, errors.Wrap(ErrQueryNode, err.Error())
	}

	return resp.Params.Convert().(QueryParamsResp), nil
}
