package oracle

import (
	"context"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/irisnet/core-sdk-go/codec"
	"github.com/irisnet/core-sdk-go/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type oracleClient struct {
	sdk.BaseClient
	codec.Codec
}

func NewClient(bc sdk.BaseClient, cdc codec.Codec) Client {
	return oracleClient{
		BaseClient: bc,
		Codec:      cdc,
	}
}

func (o oracleClient) Name() string {
	return ModuleName
}

func (o oracleClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (o oracleClient) CreateFeed(request CreateFeedRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	creator, err := o.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	msg := &MsgCreateFeed{
		FeedName:          request.FeedName,
		AggregateFunc:     request.AggregateFunc,
		ValueJsonPath:     request.ValueJsonPath,
		LatestHistory:     request.LatestHistory,
		Description:       request.Description,
		ServiceName:       request.ServiceName,
		Providers:         request.Providers,
		Input:             request.Input,
		Timeout:           request.Timeout,
		ServiceFeeCap:     request.ServiceFeeCap,
		RepeatedFrequency: request.RepeatedFrequency,
		ResponseThreshold: request.ResponseThreshold,
		Creator:           creator.String(),
	}

	return o.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (o oracleClient) StartFeed(feedName string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	creator, err := o.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	msg := &MsgStartFeed{
		FeedName: feedName,
		Creator:  creator.String(),
	}

	return o.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (o oracleClient) PauseFeed(feedName string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	creator, err := o.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	msg := &MsgPauseFeed{
		FeedName: feedName,
		Creator:  creator.String(),
	}

	return o.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (o oracleClient) EditFeedRequest(request EditFeedRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	creator, err := o.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	msg := &MsgEditFeed{
		FeedName:          request.FeedName,
		LatestHistory:     request.LatestHistory,
		Description:       request.Description,
		Providers:         request.Providers,
		Timeout:           request.Timeout,
		ServiceFeeCap:     request.ServiceFeeCap,
		RepeatedFrequency: request.RepeatedFrequency,
		ResponseThreshold: request.ResponseThreshold,
		Creator:           creator.String(),
	}

	return o.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (o oracleClient) QueryFeed(feedName string) (QueryFeedResp, error) {
	conn, err := o.GenConn()

	if err != nil {
		return QueryFeedResp{}, err
	}

	resp, err := NewQueryClient(conn).Feed(
		context.Background(),
		&QueryFeedRequest{
			FeedName: feedName,
		},
	)
	if err != nil {
		return QueryFeedResp{}, err
	}

	return resp.Feed.Convert().(QueryFeedResp), nil
}

func (o oracleClient) QueryFeeds(state string) ([]QueryFeedResp, error) {
	conn, err := o.GenConn()

	if err != nil {
		return []QueryFeedResp{}, err
	}

	resp, err := NewQueryClient(conn).Feeds(
		context.Background(),
		&QueryFeedsRequest{
			State: state,
		},
	)
	if err != nil {
		return []QueryFeedResp{}, err
	}

	return Feeds(resp.Feeds).Convert().([]QueryFeedResp), nil
}

func (o oracleClient) QueryFeedValue(feedName string) ([]QueryFeedValueResp, error) {
	conn, err := o.GenConn()

	if err != nil {
		return []QueryFeedValueResp{}, err
	}

	resp, err := NewQueryClient(conn).FeedValue(
		context.Background(),
		&QueryFeedValueRequest{
			FeedName: feedName,
		},
	)
	if err != nil {
		return []QueryFeedValueResp{}, err
	}

	return FeedValues(resp.FeedValues).Convert().([]QueryFeedValueResp), nil
}
