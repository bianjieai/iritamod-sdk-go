package integration

import (
	"fmt"
	"strings"

	"github.com/bianjieai/iritamod-sdk-go/nft"
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/stretchr/testify/require"
)

func (s IntegrationTestSuite) Test_nft() {
	baseTx := sdk.BaseTx{
		From:     s.Account().Name,
		Gas:      0,
		Memo:     "test",
		Mode:     sdk.Commit,
		Password: s.Account().Password,
	}

	denomID := strings.ToLower(s.RandStringOfLength(4))
	denomName := strings.ToLower(s.RandStringOfLength(4))
	schema := strings.ToLower(s.RandStringOfLength(10))
	issueReq := nft.IssueDenomRequest{
		ID:     denomID,
		Name:   denomName,
		Schema: schema,
	}
	res, err := s.Nft.IssueDenom(issueReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res.Hash)

	tokenID := strings.ToLower(s.RandStringOfLength(7))
	tokenName := strings.ToLower(s.RandStringOfLength(7))
	tokenData := strings.ToLower(s.RandStringOfLength(7))
	mintReq := nft.MintNFTRequest{
		Denom: denomID,
		ID:    tokenID,
		Name:  tokenName,
		URI:   fmt.Sprintf("https://%s", s.RandStringOfLength(10)),
		Data:  tokenData,
	}
	res, err = s.Nft.MintNFT(mintReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res.Hash)

	editReq := nft.EditNFTRequest{
		Denom: mintReq.Denom,
		ID:    mintReq.ID,
		URI:   fmt.Sprintf("https://%s", s.RandStringOfLength(10)),
	}
	res, err = s.Nft.EditNFT(editReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res.Hash)

	nftRes, err := s.Nft.QueryNFT(mintReq.Denom, mintReq.ID)
	require.NoError(s.T(), err)
	require.Equal(s.T(), editReq.URI, nftRes.URI)

	supply, err := s.Nft.QuerySupply(mintReq.Denom, nftRes.Creator)
	require.NoError(s.T(), err)
	require.Equal(s.T(), uint64(1), supply)

	owner, err := s.Nft.QueryOwner(nftRes.Creator, mintReq.Denom, nil)
	require.NoError(s.T(), err)
	require.Len(s.T(), owner.IDCs, 1)
	require.Len(s.T(), owner.IDCs[0].TokenIDs, 1)
	require.Equal(s.T(), tokenID, owner.IDCs[0].TokenIDs[0])

	uName := s.RandStringOfLength(10)
	pwd := "11111111"

	recipient, _, err := s.Add(uName, pwd)
	require.NoError(s.T(), err)

	transferReq := nft.TransferNFTRequest{
		Recipient: recipient,
		Denom:     mintReq.Denom,
		ID:        mintReq.ID,
		URI:       fmt.Sprintf("https://%s", s.RandStringOfLength(10)),
	}
	res, err = s.Nft.TransferNFT(transferReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res.Hash)

	owner, err = s.Nft.QueryOwner(transferReq.Recipient, mintReq.Denom, nil)
	require.NoError(s.T(), err)
	require.Len(s.T(), owner.IDCs, 1)
	require.Len(s.T(), owner.IDCs[0].TokenIDs, 1)
	require.Equal(s.T(), tokenID, owner.IDCs[0].TokenIDs[0])

	supply, err = s.Nft.QuerySupply(mintReq.Denom, transferReq.Recipient)
	require.NoError(s.T(), err)
	require.Equal(s.T(), uint64(1), supply)

	denoms, err := s.Nft.QueryDenoms(nil)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), denoms)

	d, err := s.Nft.QueryDenom(denomID)
	require.NoError(s.T(), err)
	require.Equal(s.T(), denomID, d.ID)
	require.Equal(s.T(), denomName, d.Name)
	require.Equal(s.T(), schema, d.Schema)

	col, err := s.Nft.QueryCollection(denomID, nil)
	require.NoError(s.T(), err)
	require.EqualValues(s.T(), d, col.Denom)
	require.Len(s.T(), col.NFTs, 1)
	require.Equal(s.T(), mintReq.ID, col.NFTs[0].ID)

	burnReq := nft.BurnNFTRequest{
		Denom: mintReq.Denom,
		ID:    mintReq.ID,
	}

	amount, e := sdk.ParseDecCoins("1000upoint")
	require.NoError(s.T(), e)
	_, err = s.Bank.Send(recipient, amount, baseTx)
	require.NoError(s.T(), err)

	baseTx.From = uName
	baseTx.Password = pwd
	res, err = s.Nft.BurnNFT(burnReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res.Hash)

	supply, err = s.Nft.QuerySupply(mintReq.Denom, transferReq.Recipient)
	require.NoError(s.T(), err)
	require.Equal(s.T(), uint64(0), supply)
}
