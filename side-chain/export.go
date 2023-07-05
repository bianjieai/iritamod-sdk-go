package side_chain

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	sdk "github.com/irisnet/core-sdk-go/types"
)

// Client export a group api for Admin module
type Client interface {
	sdk.Module

	CreateSpace(name, uri string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	TransferSpace(spaceId uint64, recipient string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	CreateBlockHeader(spaceId uint64, height uint64, blockHeader string, baseTx sdk.BaseTx) (sdk.ResultTx, error)

	GetSpace(spaceID uint64) (*Space, error)
	GetSpaceOfOwner(owner string, page *query.PageRequest) ([]Space, error)
	GetBlockHeader(spaceID uint64, height uint64) (string, error)
}
