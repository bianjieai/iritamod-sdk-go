package integration

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/stretchr/testify/require"

	sdk "github.com/irisnet/core-sdk-go/types"

	"github.com/bianjieai/iritamod-sdk-go/node"
)

func (s IntegrationTestSuite) TestValidator() {
	baseTx := sdk.BaseTx{
		From:          s.Account().Name,
		Password:      s.Account().Password,
		Gas:           gasWanted,
		Fee:           feeWanted,
		Mode:          sdk.Commit,
		GasAdjustment: 1.5,
	}

	cert := string(getRootPem())

	createReq := node.CreateValidatorRequest{
		Name:        "test1",
		Certificate: cert,
		Power:       10,
		Details:     "this is a test",
	}

	res1, err := s.Node.CreateValidator(createReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res1.Hash)

	validatorID, err := res1.Events.GetValue("create_validator", "validator")
	require.NoError(s.T(), err)

	res2, err := s.Node.QueryValidator(validatorID)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res2)

	res3, err := s.Node.QueryValidators(nil)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res3)

	updateReq := node.UpdateValidatorRequest{
		ID:          validatorID,
		Name:        "test2",
		Certificate: cert,
		Power:       10,
		Details:     "this is a updated test",
	}
	res4, err := s.Node.UpdateValidator(updateReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res4.Hash)

	res5, err := s.Node.QueryValidator(validatorID)
	require.NoError(s.T(), err)
	require.Equal(s.T(), updateReq.Name, res5.Name)
	require.Equal(s.T(), updateReq.Details, res5.Details)

	res6, err := s.Node.RemoveValidator(validatorID, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res6.Hash)

	res7, err := s.Node.QueryValidator(validatorID)
	require.Error(s.T(), err)
	require.Empty(s.T(), res7.Name)

	grantNodeReq := node.GrantNodeRequest{
		Name:        "test3",
		Certificate: cert,
		Details:     "this is a grantNode test",
	}
	res8, err := s.Node.GrantNode(grantNodeReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res8.Hash)

	noid, e := res8.Events.GetValue("grant_node", "id")
	require.NoError(s.T(), e)

	res9, err := s.Node.QueryNode(noid)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res9)

	ns, err := s.Node.QueryNodes(nil)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), ns)

	res10, err := s.Node.RevokeNode(noid, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res10.Hash)

}

func getRootPem() []byte {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = filepath.Dir(path)
	path = filepath.Join(path, "integration_test/scripts/root_cert.pem")
	bz, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bz
}
