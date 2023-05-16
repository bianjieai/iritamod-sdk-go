package layer2

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/query"
)

// Client export a group api for Admin module
type Client interface {
	sdk.Module

	CreateL2Space(baseTx sdk.BaseTx) (sdk.ResultTx, error)
	TransferL2Space(spaceId uint64, recipient string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	CreateL2Record(spaceId uint64, height uint64, blockHeader string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	CreateNFTs(spaceId uint64, classId string, nfts []*TokenForNFT, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	UpdateNFTs(spaceId uint64, classId string, nfts []*TokenForNFT, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	DeleteNFTs(spaceId uint64, classId string, nftIds []string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	DepositClassForNFT(spaceId uint64, classId string, baseURI string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	WithdrawClassForNFT(spaceId uint64, classId string, owner string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	DepositTokenForNFT(spaceId uint64, classId string, nftId string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	WithdrawTokenForNFT(spaceId uint64, classId string, nftId string, owner string, baseTx sdk.BaseTx) (sdk.ResultTx, error)

	GetSpace(spaceID uint64) (*Space, error)
	GetSpaceOfOwner(owner string) ([]uint64, error)
	GetRecord(spaceID uint64, height uint64) (string, error)
	GetClassForNFT(classID string) (*ClassForNFT, error)
	GetClassesForNFT(page *query.PageRequest) ([]*ClassForNFT, error)
	GetTokenForNFT(spaceID uint64, classID string, nftID string) (string, error)
	GetTokensOfOwnerForNFT(spaceID uint64, classID string, owner string, page *query.PageRequest) ([]string, error)
	GetCollectionForNFT(spaceID uint64, classID string, page *query.PageRequest) ([]*TokenForNFT, error)
}
