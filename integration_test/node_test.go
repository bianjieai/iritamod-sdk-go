package integration

import (
	"github.com/stretchr/testify/require"

	"github.com/bianjieai/irita-sdk-go/node"
	sdk "github.com/irisnet/core-sdk-go/types"
)

func (s IntegrationTestSuite) TestValidator() {
	baseTx := sdk.BaseTx{
		From:     s.Account().Name,
		Gas:      0,
		Memo:     "test",
		Mode:     sdk.Commit,
		Password: s.Account().Password,
	}

	cert := `-----BEGIN CERTIFICATE-----
MIICBjCCAaugAwIBAgIUZFNJU4ANpSzL6ggat80j+h6gnDQwCgYIKoEcz1UBg3Uw
WDELMAkGA1UEBhMCQ04xDTALBgNVBAgMBHJvb3QxDTALBgNVBAcMBHJvb3QxDTAL
BgNVBAoMBHJvb3QxDTALBgNVBAsMBHJvb3QxDTALBgNVBAMMBHJvb3QwHhcNMjEw
OTE3MDMzNzAwWhcNMjIwOTE3MDMzNzAwWjBYMQswCQYDVQQGEwJDTjENMAsGA1UE
CAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9vdDENMAsGA1UECwwE
cm9vdDENMAsGA1UEAwwEcm9vdDBZMBMGByqGSM49AgEGCCqBHM9VAYItA0IABA8G
cAky8uGAgeDdJ7sbMj3VioSHXCklHu4Lck3hglspgVnI7kjJR+rExhnqhWO4u4a1
qM6y18Z9SxvnLRNzAO6jUzBRMB0GA1UdDgQWBBROrZ7qRryLYFsRVe71xLw16tJY
tzAfBgNVHSMEGDAWgBROrZ7qRryLYFsRVe71xLw16tJYtzAPBgNVHRMBAf8EBTAD
AQH/MAoGCCqBHM9VAYN1A0kAMEYCIQDDGh24GkPzaA5jlq5g2TRENpXhenDDNIgi
MY5maJdfuQIhAOUcqtg2W2+gONf7MCpFCcovPNNwXc+rI0WaGPrdJu8k
-----END CERTIFICATE-----`

	createReq := node.CreateValidatorRequest{
		Name:        "test1",
		Certificate: cert,
		Power:       10,
		Details:     "this is a test",
	}

	rs, err := s.Node.CreateValidator(createReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

	validatorID, er := rs.Events.GetValue("create_validator", "validator")
	require.NoError(s.T(), er)

	v, err := s.Node.QueryValidator(validatorID)
	require.NoError(s.T(), err)

	vs, err := s.Node.QueryValidators(nil)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), vs)

	updateReq := node.UpdateValidatorRequest{
		ID:          validatorID,
		Name:        "test2",
		Certificate: cert,
		Power:       10,
		Details:     "this is a updated test",
	}
	rs, err = s.Node.UpdateValidator(updateReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

	v, err = s.Node.QueryValidator(validatorID)
	require.NoError(s.T(), err)
	require.Equal(s.T(), updateReq.Name, v.Name)
	require.Equal(s.T(), updateReq.Details, v.Details)

	rs, err = s.Node.RemoveValidator(validatorID, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

	v, err = s.Node.QueryValidator(validatorID)
	require.Error(s.T(), err)

	grantNodeReq := node.GrantNodeRequest{
		Name:        "test3",
		Certificate: cert,
		Details:     "this is a grantNode test",
	}
	rs, err = s.Node.GrantNode(grantNodeReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

	noid, e := rs.Events.GetValue("grant_node", "id")
	require.NoError(s.T(), e)

	n, err := s.Node.QueryNode(noid)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), n)

	ns, err := s.Node.QueryNodes(nil)
	require.NoError(s.T(), err)
	require.Equal(s.T(), 2, len(ns))

	rs, err = s.Node.RevokeNode(noid, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

}
