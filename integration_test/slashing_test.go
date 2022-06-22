package integration

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/stretchr/testify/require"
)

func (s IntegrationTestSuite) TestUnjailValidator() {
	baseTx := sdk.BaseTx{
		From:          s.Account().Name,
		Password:      s.Account().Password,
		Gas:           gasWanted,
		Fee:           feeWanted,
		Mode:          sdk.Commit,
		GasAdjustment: 1.5,
	}

	id := "iaa1d3p96vyvekh8wwrse5s2y7vhq9n6967mefsrav"

	// TODO: this is a bad test
	res, err := s.Slashing.UnjailValidator(id, baseTx)
	require.Error(s.T(), err)
	require.Empty(s.T(), res.Hash)
}
