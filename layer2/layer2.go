package layer2

import (
	"context"

	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/query"
)

type layer2Client struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) Client {
	return layer2Client{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (l layer2Client) Name() string {
	return ModuleName
}

func (l layer2Client) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (l layer2Client) CreateL2Space(name, uri string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := l.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgCreateL2Space{
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

func (l layer2Client) TransferL2Space(spaceId uint64, recipient string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := l.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgTransferL2Space{
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

func (l layer2Client) CreateL2BlockHeader(spaceId uint64, height uint64, blockHeader string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := l.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgCreateL2BlockHeader{
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

func (l layer2Client) CreateNFTs(spaceId uint64, classId string, tokens []TokenForNFT, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := l.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgCreateNFTs{
		Sender:  sender.String(),
		SpaceId: spaceId,
		ClassId: classId,
		Tokens:  tokens,
	}
	resultTx, err := l.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return resultTx, nil
}

func (l layer2Client) UpdateNFTs(spaceId uint64, classId string, tokens []TokenForNFT, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := l.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgUpdateNFTs{
		Sender:  sender.String(),
		SpaceId: spaceId,
		ClassId: classId,
		Tokens:  tokens,
	}
	resultTx, err := l.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return resultTx, nil
}

func (l layer2Client) DeleteNFTs(spaceId uint64, classId string, tokenIds []string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := l.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgDeleteNFTs{
		Sender:   sender.String(),
		SpaceId:  spaceId,
		ClassId:  classId,
		TokenIds: tokenIds,
	}
	resultTx, err := l.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return resultTx, nil
}

func (l layer2Client) UpdateClassesForNFT(spaceId uint64, ClassUpdatesForNft []UpdateClassForNFT, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := l.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}
	msg := &MsgUpdateClassesForNFT{
		Sender:             sender.String(),
		SpaceId:            spaceId,
		ClassUpdatesForNft: ClassUpdatesForNft,
	}
	resultTx, err := l.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return resultTx, nil
}

func (l layer2Client) DepositClassForNFT(spaceId uint64, classId string, baseURI string, recipient string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := l.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgDepositClassForNFT{
		SpaceId:   spaceId,
		ClassId:   classId,
		BaseUri:   baseURI,
		Recipient: recipient,
		Sender:    sender.String(),
	}
	resultTx, err := l.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return resultTx, nil
}

func (l layer2Client) WithdrawClassForNFT(spaceId uint64, classId string, owner string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := l.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgWithdrawClassForNFT{
		SpaceId: spaceId,
		Sender:  sender.String(),
		ClassId: classId,
		Owner:   owner,
	}
	resultTx, err := l.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return resultTx, nil
}

func (l layer2Client) DepositTokenForNFT(spaceId uint64, classId string, tokenId string, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := l.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgDepositTokenForNFT{
		Sender:  sender.String(),
		ClassId: classId,
		TokenId: tokenId,
		SpaceId: spaceId,
	}
	resultTx, err := l.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return resultTx, nil
}

func (l layer2Client) WithdrawTokenForNFT(
	spaceId uint64,
	classId, tokenId, owner, name, uri, uriHash, data string,
	baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := l.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrQueryAddress, err.Error())
	}

	msg := &MsgWithdrawTokenForNFT{
		SpaceId: spaceId,
		ClassId: classId,
		TokenId: tokenId,
		Owner:   owner,
		Name:    name,
		Uri:     uri,
		UriHash: uriHash,
		Data:    data,
		Sender:  sender.String(),
	}
	resultTx, err := l.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, sdk.WrapWithMessage(ErrBuildAndSend, err.Error())
	}
	return resultTx, nil
}

// Query Function

// GetSpace returns the space info of the given spaceID
func (l layer2Client) GetSpace(spaceID uint64) (*Space, error) {
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
func (l layer2Client) GetSpaceOfOwner(owner string, page *query.PageRequest) ([]Space, error) {
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

func (l layer2Client) GetL2BlockHeader(spaceID uint64, height uint64) (string, error) {
	conn, err := l.GenConn()
	if err != nil {
		return "", sdk.WrapWithMessage(ErrGenConn, err.Error())
	}
	resp, err := NewQueryClient(conn).L2BlockHeader(
		context.Background(),
		&QueryL2BlockHeaderRequest{
			SpaceId: spaceID,
			Height:  height,
		})
	if err != nil {
		return "", sdk.WrapWithMessage(ErrQueryPerm, err.Error())
	}
	return resp.Header, nil
}

func (l layer2Client) GetClassForNFT(classID string) (*ClassForNFT, error) {
	conn, err := l.GenConn()
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}
	resp, err := NewQueryClient(conn).ClassForNFT(
		context.Background(),
		&QueryClassForNFTRequest{
			ClassId: classID,
		})
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrQueryPerm, err.Error())
	}
	return resp.Class, nil
}

func (l layer2Client) GetClassesForNFT(page *query.PageRequest) ([]ClassForNFT, error) {
	conn, err := l.GenConn()
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}
	resp, err := NewQueryClient(conn).ClassesForNFT(
		context.Background(),
		&QueryClassesForNFTRequest{
			Pagination: page,
		})
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrQueryPerm, err.Error())
	}
	return resp.Classes, nil
}

func (l layer2Client) GetTokenForNFT(spaceID uint64, classID string, tokenID string) (string, error) {
	conn, err := l.GenConn()
	if err != nil {
		return "", sdk.WrapWithMessage(ErrGenConn, err.Error())
	}
	resp, err := NewQueryClient(conn).TokenForNFT(
		context.Background(),
		&QueryTokenForNFTRequest{
			SpaceId: spaceID,
			ClassId: classID,
			TokenId: tokenID,
		})
	if err != nil {
		return "", sdk.WrapWithMessage(ErrQueryPerm, err.Error())
	}
	return resp.Owner, nil
}

func (l layer2Client) GetCollectionForNFT(spaceID uint64, classID string, page *query.PageRequest) ([]TokenForNFT, error) {
	conn, err := l.GenConn()
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}
	resp, err := NewQueryClient(conn).CollectionForNFT(
		context.Background(),
		&QueryCollectionForNFTRequest{
			SpaceId:    spaceID,
			ClassId:    classID,
			Pagination: page,
		})
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrQueryPerm, err.Error())
	}
	return resp.Tokens, nil
}

func (l layer2Client) GetTokensOfOwnerForNFT(spaceID uint64, classID string, owner string, page *query.PageRequest) ([]TokenForNFTByOwner, error) {
	conn, err := l.GenConn()
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrGenConn, err.Error())
	}
	resp, err := NewQueryClient(conn).TokensOfOwnerForNFT(
		context.Background(),
		&QueryTokensOfOwnerForNFTRequest{
			SpaceId:    spaceID,
			ClassId:    classID,
			Owner:      owner,
			Pagination: page,
		})
	if err != nil {
		return nil, sdk.WrapWithMessage(ErrQueryPerm, err.Error())
	}
	return resp.Tokens, nil
}

func (l layer2Client) GetBaseUriForNFT(classId string) (string, error) {
	conn, err := l.GenConn()
	if err != nil {
		return "", sdk.WrapWithMessage(ErrGenConn, err.Error())
	}
	resp, err := NewQueryClient(conn).BaseUriForNFT(
		context.Background(),
		&QueryBaseUriForNFTRequest{
			ClassId: classId,
		})
	if err != nil {
		return "", sdk.WrapWithMessage(ErrQueryPerm, err.Error())
	}
	return resp.BaseUri, nil
}
func (l layer2Client) GetTokenUriForNFT(spaceId uint64, classId string, tokenId string) (string, error) {
	conn, err := l.GenConn()
	if err != nil {
		return "", sdk.WrapWithMessage(ErrGenConn, err.Error())
	}
	resp, err := NewQueryClient(conn).TokenUriForNFT(
		context.Background(),
		&QueryTokenUriForNFTRequest{
			SpaceId: spaceId,
			ClassId: classId,
			TokenId: tokenId,
		})
	if err != nil {
		return "", sdk.WrapWithMessage(ErrQueryPerm, err.Error())
	}
	return resp.TokenUri, nil
}
