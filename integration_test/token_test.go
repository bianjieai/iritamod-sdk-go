package integration

import (
	"strings"

	"github.com/bianjieai/iritamod-sdk-go/params"

	"github.com/stretchr/testify/require"

	"github.com/bianjieai/iritamod-sdk-go/token"
	sdk "github.com/irisnet/core-sdk-go/types"
)

func (s IntegrationTestSuite) TestToken() {
	baseTx := sdk.BaseTx{
		From:     s.Account().Name,
		Gas:      200000,
		Memo:     "test",
		Mode:     sdk.Commit,
		Password: s.Account().Password,
	}

	issueTokenReq := token.IssueTokenRequest{
		Symbol:        strings.ToLower(s.RandStringOfLength(3)),
		Name:          s.RandStringOfLength(8),
		Scale:         9,
		MinUnit:       strings.ToLower(s.RandStringOfLength(3)),
		InitialSupply: 10000000,
		MaxSupply:     21000000,
		Mintable:      true,
	}
	var request1 = []params.UpdateParamRequest{{
		Module: "token",
		Key:    "IssueTokenBaseFee",
		Value:  `{"denom":"point","amount":"20"}`,
	}}

	rs1, err := s.Params.UpdateParams(request1, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs1.Hash)

	//test issue token
	rs, err := s.Token.IssueToken(issueTokenReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

	//test mint token
	receipt := s.GetRandAccount().Address.String()
	rs, err = s.Token.MintToken(issueTokenReq.Symbol, 1000, receipt, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

	account, err := s.Bank.QueryAccount(receipt)
	require.NoError(s.T(), err)

	amt := sdk.NewIntWithDecimal(1000, int(issueTokenReq.Scale))
	require.Equal(s.T(), amt, account.Coins.AmountOf(issueTokenReq.MinUnit))

	editTokenReq := token.EditTokenRequest{
		Symbol:    issueTokenReq.Symbol,
		Name:      "ethereum network",
		MaxSupply: 20000000,
		Mintable:  false,
	}

	//test edit token
	rs, err = s.Token.EditToken(editTokenReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

	//test transfer token
	rs, err = s.Token.TransferToken(receipt, issueTokenReq.Symbol, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

	t1, er := s.Token.QueryToken(issueTokenReq.Symbol)
	require.NoError(s.T(), er)
	require.Equal(s.T(), t1.Name, editTokenReq.Name)
	require.Equal(s.T(), t1.MaxSupply, editTokenReq.MaxSupply)
	require.Equal(s.T(), t1.Mintable, editTokenReq.Mintable)
	require.Equal(s.T(), receipt, t1.Owner)

	tokens, er := s.Token.QueryTokens(t1.Owner, nil)
	require.NoError(s.T(), er)
	require.Contains(s.T(), tokens, t1)

	feeToken, er := s.Token.QueryFees(issueTokenReq.Symbol)
	require.NoError(s.T(), er)
	require.Equal(s.T(), true, feeToken.Exist)
	require.Equal(s.T(), "20000000upoint", feeToken.IssueFee.String())
	require.Equal(s.T(), "2000000upoint", feeToken.MintFee.String())

	res, er := s.Token.QueryParams()
	require.NoError(s.T(), er)
	require.Equal(s.T(), "0.100000000000000000", res.MintTokenFeeRatio)
	require.Equal(s.T(), "0.400000000000000000", res.TokenTaxRate)
	require.Equal(s.T(), "20point", res.IssueTokenBaseFee)
}
