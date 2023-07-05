package side_chain

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type sideChainClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) Client {
	return sideChainClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (l sideChainClient) Name() string {
	return ModuleName
}

func (l sideChainClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (l sideChainClient) CreateSpace(name, uri string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := l.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgCreateSpace{
		Sender: sender.String(),
		Name:   name,
		Uri:    uri,
	}
	resultTx, err := l.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return resultTx, nil
}

func (l sideChainClient) TransferSpace(spaceId uint64, recipient string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := l.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgTransferSpace{
		Sender:    sender.String(),
		SpaceId:   spaceId,
		Recipient: recipient,
	}
	resultTx, err := l.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return resultTx, nil
}

func (l sideChainClient) CreateBlockHeader(spaceId uint64, height uint64, blockHeader string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := l.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgCreateBlockHeader{
		Sender:  sender.String(),
		SpaceId: spaceId,
		Height:  height,
		Header:  blockHeader,
	}
	resultTx, err := l.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return resultTx, nil
}

// Query Function

// GetSpace returns the space info of the given spaceID
func (l sideChainClient) GetSpace(spaceID uint64) (*Space, error) {
	conn, err := l.GenConn()
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}
	resp, err := NewQueryClient(conn).Space(
		context.Background(),
		&QuerySpaceRequest{SpaceId: spaceID})
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrQueryPerm, err.Error())
	}
	return resp.Space, nil
}

// GetSpaceOfOwner returns all spaces
func (l sideChainClient) GetSpaceOfOwner(owner string, page *query.PageRequest) ([]Space, error) {
	conn, err := l.GenConn()
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}
	resp, err := NewQueryClient(conn).SpaceOfOwner(
		context.Background(),
		&QuerySpaceOfOwnerRequest{
			Owner:      owner,
			Pagination: page,
		})
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrQueryPerm, err.Error())
	}
	return resp.Spaces, nil
}

func (l sideChainClient) GetBlockHeader(spaceID uint64, height uint64) (string, error) {
	conn, err := l.GenConn()
	if err != nil {
		return "", sdk.WrapWithMessage(ErrGenConn, err.Error())
	}
	resp, err := NewQueryClient(conn).BlockHeader(
		context.Background(),
		&QueryBlockHeaderRequest{
			SpaceId: spaceID,
			Height:  height,
		})
	if err != nil {
		return "", sdk.WrapWithMessage(ErrQueryPerm, err.Error())
	}
	return resp.Header, nil
}
