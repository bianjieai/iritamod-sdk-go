package integration

import (
	"github.com/bianjieai/iritamod-sdk-go/perm"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/stretchr/testify/require"
)

func (s IntegrationTestSuite) TestLeayer2() {
	baseTx := types.BaseTx{
		From:          s.Account().Name,
		Password:      s.Account().Password,
		Gas:           gasWanted,
		Fee:           feeWanted,
		Mode:          types.Commit,
		GasAdjustment: 1.5,
	}

	acc := s.GetRandAccount()
	// 1. Create Space For Chain
	// user role is LAYER2_USER

	// if user want to create space for chain, user must have LAYER2_USER role
	roles := []perm.Role{
		perm.RoleLayer2User,
	}
	_, err := s.Perm.AssignRoles(acc.Address.String(), roles, baseTx)
	require.NoError(s.T(), err)
	spaceName := ""
	spaceURI := ""
	_, err = s.SideChainClient.CreateSpace(spaceName, spaceURI, baseTx)
	require.NoError(s.T(), err)
	var exceptSpaceId uint64 = 1
	space, err := s.SideChainClient.GetSpace(exceptSpaceId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), space.Name, spaceName)
	require.Equal(s.T(), space.Uri, spaceURI)

}
