package oracle

import (
	"time"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/bianjieai/iritamod-sdk-go/service"
	sdk "github.com/irisnet/core-sdk-go/types"
	types "github.com/irisnet/core-sdk-go/types"
)

var (
	_ Client = oracleClient{}
)

// expose Oracle module api for user
type Client interface {
	sdk.Module

	CreateFeed(request CreateFeedRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error)
	StartFeed(feedName string, baseTx sdk.BaseTx) (ctypes.ResultTx, error)
	PauseFeed(feedName string, baseTx sdk.BaseTx) (ctypes.ResultTx, error)
	EditFeedRequest(request EditFeedRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error)

	QueryFeed(feedName string) (QueryFeedResp, error)
	QueryFeeds(state string) ([]QueryFeedResp, error)
	QueryFeedValue(feedName string) ([]QueryFeedValueResp, error)
}

type CreateFeedRequest struct {
	FeedName          string
	LatestHistory     uint64
	Description       string
	ServiceName       string
	Providers         []string
	Input             string
	Timeout           int64
	ServiceFeeCap     []types.Coin
	RepeatedFrequency uint64
	AggregateFunc     string
	ValueJsonPath     string
	ResponseThreshold uint32
}

type EditFeedRequest struct {
	FeedName          string
	Description       string
	LatestHistory     uint64
	Providers         []string
	Timeout           int64
	ServiceFeeCap     []types.Coin
	RepeatedFrequency uint64
	ResponseThreshold uint32
}

type QueryFeedResp struct {
	Feed              *Feed
	ServiceName       string
	Providers         []string
	Input             string
	Timeout           int64
	ServiceFeeCap     []types.Coin
	RepeatedFrequency uint64
	ResponseThreshold uint32
	State             service.RequestContextState
}

type QueryFeedValueResp struct {
	Data      string
	Timestamp time.Time
}
