package integration

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/stretchr/testify/require"

	"github.com/bianjieai/iritamod-sdk-go/node"
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

	//	cert := `-----BEGIN CERTIFICATE-----
	//MIICBjCCAaugAwIBAgIUGaV9lOO9McFCQ8rUnVr13G1LIM4wCgYIKoEcz1UBg3Uw
	//WDELMAkGA1UEBhMCQ04xDTALBgNVBAgMBHJvb3QxDTALBgNVBAcMBHJvb3QxDTAL
	//BgNVBAoMBHJvb3QxDTALBgNVBAsMBHJvb3QxDTALBgNVBAMMBHJvb3QwHhcNMjEw
	//OTIyMDM1NzA3WhcNMjIwOTIyMDM1NzA3WjBYMQswCQYDVQQGEwJDTjENMAsGA1UE
	//CAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9vdDENMAsGA1UECwwE
	//cm9vdDENMAsGA1UEAwwEcm9vdDBZMBMGByqGSM49AgEGCCqBHM9VAYItA0IABDpd
	//c1KOPowpvC9YCDDTYp/MYcoDSawUHoQO8Dl+yQzyqeWA3Gko3dosF9l2rM5gUHp6
	//YS/hMhtAhjFnPLm0GfijUzBRMB0GA1UdDgQWBBSW7EMp99BmeVIUiGZ2yrI1AW8B
	//rTAfBgNVHSMEGDAWgBSW7EMp99BmeVIUiGZ2yrI1AW8BrTAPBgNVHRMBAf8EBTAD
	//AQH/MAoGCCqBHM9VAYN1A0kAMEYCIQD8cqMmzHaK1+idd8dN+MUQpU3+N6gcfscS
	//9HbFEl+IQAIhAJCZ2rQVhfit8Hif02Dic9jky8BWvA7UxyK109mVKH86
	//-----END CERTIFICATE-----`

	cert := string(getRootPem())

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
	require.NotEmpty(s.T(), ns)

	rs, err = s.Node.RevokeNode(noid, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

}

func getRootPem() []byte {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = filepath.Dir(path)
	path = filepath.Join(path, "integration_test/scripts/testnet/root_cert.pem")
	bz, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bz
}
