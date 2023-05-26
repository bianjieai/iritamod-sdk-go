package layer2

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/query"
)

// Client export a group api for Admin module
type Client interface {
	sdk.Module

	CreateL2Space(name, uri string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	TransferL2Space(spaceId uint64, recipient string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	CreateNFTs(spaceId uint64, classId string, tokens []TokenForNFT, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	UpdateNFTs(spaceId uint64, classId string, tokens []TokenForNFT, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	DeleteNFTs(spaceId uint64, classId string, tokenIds []string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	UpdateClassesForNFT(spaceId uint64, ClassUpdatesForNft []UpdateClassForNFT, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	DepositClassForNFT(spaceId uint64, classId string, baseURI string, recipient string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	WithdrawClassForNFT(spaceId uint64, classId string, owner string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	DepositTokenForNFT(spaceId uint64, classId string, tokenId string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	WithdrawTokenForNFT(spaceId uint64, classId, tokenId, owner, name, uri, uriHash, data string, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	CreateL2BlockHeader(spaceId uint64, height uint64, blockHeader string, baseTx sdk.BaseTx) (sdk.ResultTx, error)

	GetSpace(spaceID uint64) (*Space, error)
	GetSpaceOfOwner(owner string, page *query.PageRequest) ([]Space, error)
	GetL2BlockHeader(spaceID uint64, height uint64) (string, error)
	GetClassForNFT(classID string) (*ClassForNFT, error)
	GetClassesForNFT(page *query.PageRequest) ([]ClassForNFT, error)
	GetTokenForNFT(spaceID uint64, classID string, nftID string) (string, error)
	GetTokensOfOwnerForNFT(spaceID uint64, classID string, owner string, page *query.PageRequest) ([]TokenForNFTByOwner, error)
	GetCollectionForNFT(spaceID uint64, classID string, page *query.PageRequest) ([]TokenForNFT, error)
	GetBaseUriForNFT(classId string) (string, error)
	GetTokenUriForNFT(spaceId uint64, classId string, tokenId string) (string, error)
}
