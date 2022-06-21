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

	request := []params.UpdateParamRequest{
		{
			Module: "node",
			Key:    "HistoricalEntries",
			Value:  "110",
		},
	}

	res, err := s.Params.UpdateParams(request, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res.Hash)

	// As no query method, we can only check tx hash
}
