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
		Name:        "create",
		Certificate: cert,
		Power:       10,
		Details:     "this is a create test",
	}

	// create validator
	createResp, err := s.Node.CreateValidator(createReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), createResp.Hash)

	vId, err := createResp.Events.GetValue("create_validator", "validator")
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), vId)
	println(vId)

	queryResp1, err := s.Node.QueryValidator(vId)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), queryResp1)
	require.Equal(s.T(), queryResp1.Details, createReq.Details)

	queryResp2, err := s.Node.QueryValidators(nil)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), queryResp2)
	require.Contains(s.T(), queryResp2, queryResp1)

	// update validator
	updateReq := node.UpdateValidatorRequest{
		ID:          vId,
		Name:        "update",
		Certificate: cert,
		Power:       15,
		Details:     "this is an update test",
	}
	updateResp, err := s.Node.UpdateValidator(updateReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), updateResp.Hash)

	queryResp3, err := s.Node.QueryValidator(vId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), updateReq.Name, queryResp3.Name)
	require.Equal(s.T(), updateReq.Details, queryResp3.Details)
	require.Equal(s.T(), updateReq.Power, queryResp3.Power)

	// remove validator
	removeResp, err := s.Node.RemoveValidator(vId, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), removeResp.Hash)

	queryResp4, err := s.Node.QueryValidator(vId)
	require.Error(s.T(), err)
	require.Empty(s.T(), queryResp4.Name)
}

func (s IntegrationTestSuite) TestGrantRevoke() {
	baseTx := sdk.BaseTx{
		From:          s.Account().Name,
		Password:      s.Account().Password,
		Gas:           gasWanted,
		Fee:           feeWanted,
		Mode:          sdk.Commit,
		GasAdjustment: 1.5,
	}

	cert := string(getRootPem())

	// grant node
	grantNodeReq := node.GrantNodeRequest{
		Name:        "grant",
		Certificate: cert,
		Details:     "this is a grant test",
	}

	grantResp, err := s.Node.GrantNode(grantNodeReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), grantResp.Hash)

	nodeId, err := grantResp.Events.GetValue("grant_node", "id")
	require.NoError(s.T(), err)

	queryResp1, err := s.Node.QueryNode(nodeId)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), queryResp1)

	queryResp2, err := s.Node.QueryNodes(nil)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), queryResp2)

	// revoke node
	revokeResp, err := s.Node.RevokeNode(nodeId, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), revokeResp.Hash)

	queryResp3, err := s.Node.QueryNode(nodeId)
	require.Error(s.T(), err)
	require.Empty(s.T(), queryResp3)
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
