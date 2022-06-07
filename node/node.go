package node

import (
	"context"

	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
	query "github.com/irisnet/core-sdk-go/types/query"
)

type nodeClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) Client {
	return nodeClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (n nodeClient) Name() string {
	return ModuleName
}

func (n nodeClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (n nodeClient) CreateValidator(request CreateValidatorRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	creator, err := n.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
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
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (n nodeClient) UpdateValidator(request UpdateValidatorRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	creator, err := n.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	vID, er := sdk.HexBytesFrom(request.ID)
	if er != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrHex, err.Error())
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
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (n nodeClient) RemoveValidator(id string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	creator, err := n.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	vID, er := sdk.HexBytesFrom(id)
	if er != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrHex, err.Error())
	}
	msg := &MsgRemoveValidator{
		Id:       vID.String(),
		Operator: creator.String(),
	}

	send, err := n.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (n nodeClient) GrantNode(request GrantNodeRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	creator, err := n.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgGrantNode{
		Name:        request.Name,
		Certificate: request.Certificate,
		Operator:    creator.String(),
	}

	send, err := n.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (n nodeClient) RevokeNode(nodeId string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	creator, err := n.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	vID, er := sdk.HexBytesFrom(nodeId)
	if er != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrHex, er.Error())
	}

	msg := &MsgRevokeNode{
		Id:       vID.String(),
		Operator: creator.String(),
	}

	send, err := n.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (n nodeClient) QueryValidators(pageReq *query.PageRequest) ([]QueryValidatorResp, error) {
	conn, err := n.GenConn()

	if err != nil {
		return nil, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}

	resp, err := NewQueryClient(conn).Validators(
		context.Background(),
		&QueryValidatorsRequest{
			Pagination: pageReq,
		},
	)
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrQueryNode, err.Error())
	}

	return validators(resp.Validators).Convert().([]QueryValidatorResp), nil
}

func (n nodeClient) QueryValidator(id string) (QueryValidatorResp, error) {
	conn, err := n.GenConn()

	if err != nil {
		return QueryValidatorResp{}, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}

	vID, err := sdk.HexBytesFrom(id)
	if err != nil {
		return QueryValidatorResp{}, sdk.WrapWithMessage(ErrHex, err.Error())
	}

	resp, err := NewQueryClient(conn).Validator(
		context.Background(),
		&QueryValidatorRequest{
			Id: vID.String(),
		},
	)
	if err != nil {
		return QueryValidatorResp{}, sdk.WrapWithMessage(ErrQueryNode, err.Error())
	}

	return resp.Validator.Convert().(QueryValidatorResp), nil
}

func (n nodeClient) QueryNodes(pageReq *query.PageRequest) ([]QueryNodeResp, error) {
	conn, err := n.GenConn()

	if err != nil {
		return nil, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}

	resp, err := NewQueryClient(conn).Nodes(
		context.Background(),
		&QueryNodesRequest{
			Pagination: pageReq,
		},
	)
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrQueryNode, err.Error())
	}

	return nodes(resp.Nodes).Convert().([]QueryNodeResp), nil
}

func (n nodeClient) QueryNode(id string) (QueryNodeResp, error) {
	conn, err := n.GenConn()

	if err != nil {
		return QueryNodeResp{}, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}

	vID, err := sdk.HexBytesFrom(id)
	if err != nil {
		return QueryNodeResp{}, sdk.WrapWithMessage(ErrHex, err.Error())
	}

	resp, err := NewQueryClient(conn).Node(
		context.Background(),
		&QueryNodeRequest{
			Id: vID.String(),
		},
	)
	if err != nil {
		return QueryNodeResp{}, sdk.WrapWithMessage(ErrQueryNode, err.Error())
	}

	return resp.Node.Convert().(QueryNodeResp), nil
}

func (n nodeClient) QueryParams() (QueryParamsResp, error) {
	conn, err := n.GenConn()

	if err != nil {
		return QueryParamsResp{}, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}

	resp, err := NewQueryClient(conn).Params(
		context.Background(),
		&QueryParamsRequest{},
	)
	if err != nil {
		return QueryParamsResp{}, sdk.WrapWithMessage(ErrQueryNode, err.Error())
	}

	return resp.Params.Convert().(QueryParamsResp), nil
}
