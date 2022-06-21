package integration

import (
	"github.com/bianjieai/iritamod-sdk-go/params"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/stretchr/testify/require"
)

func (s IntegrationTestSuite) TestUpdateParams() {
	baseTx := types.BaseTx{
		From:          s.Account().Name,
		Password:      s.Account().Password,
		Gas:           gasWanted,
		Fee:           feeWanted,
		Mode:          types.Commit,
		GasAdjustment: 1.5,
	}

	req1 := []params.UpdateParamRequest{
		{
			Module: "node",
			Key:    "HistoricalEntries",
			Value:  "110",
		},
	}

	// success tx
	res, err := s.Params.UpdateParams(req1, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res.Hash)

	req2 := []params.UpdateParamRequest{
		{
			Module: "node",
			Key:    "",
			Value:  "110",
		},
	}

	// failed tx
	res2, err := s.Params.UpdateParams(req2, baseTx)
	require.Error(s.T(), err)
	require.Empty(s.T(), res2.Hash)
}
