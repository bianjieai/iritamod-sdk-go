package integration

import (
	"context"
	"encoding/hex"
	"strings"

	iritaidentity "github.com/bianjieai/iritamod-sdk-go/identity"

	"github.com/irisnet/core-sdk-go/common/uuid"
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/sm2"
)

func (s IntegrationTestSuite) Test_GetStatus() {
	status, err := s.Client.Status(context.Background())
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), status)
}

func (s IntegrationTestSuite) Test_Identity() {
	baseTx := sdk.BaseTx{
		From:     s.Account().Name,
		Gas:      0,
		Memo:     "test",
		Mode:     sdk.Commit,
		Password: s.Account().Password,
	}

	uuidGenerator, _ := uuid.NewV4()
	id := sdk.HexStringFrom(uuidGenerator.Bytes())

	testPubKeySM2 := sm2.GenPrivKey().PubKeySm2()
	testCredentials := "https://kyc.com/user/10001"
	testCertificate := ""

	pubKeyInfo := iritaidentity.PubkeyInfo{
		PubKey:     strings.ToUpper(hex.EncodeToString(testPubKeySM2[:])),
		PubKeyAlgo: iritaidentity.SM2,
	}
	request := iritaidentity.CreateIdentityRequest{
		ID:          id,
		PubkeyInfo:  &pubKeyInfo,
		Certificate: testCertificate,
		Credentials: &testCredentials,
	}

	rest, err := s.Identity.QueryIdentity(id)
	require.Empty(s.T(), rest)
	require.Error(s.T(), err)

	rs, err := s.Identity.CreateIdentity(request, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

	res, err := s.Identity.QueryIdentity(id)
	require.NoError(s.T(), err)
	require.Equal(s.T(), res.Credentials, testCredentials)
	require.Contains(s.T(), res.PubkeyInfos, pubKeyInfo)

	test2PubKeySM2 := sm2.GenPrivKey().PubKeySm2()
	pubKeyInfo2 := iritaidentity.PubkeyInfo{
		PubKey:     sdk.HexStringFrom(test2PubKeySM2[:]),
		PubKeyAlgo: iritaidentity.SM2,
	}

	req2 := iritaidentity.UpdateIdentityRequest{
		ID:          id,
		PubkeyInfo:  &pubKeyInfo2,
		Certificate: testCertificate,
		Credentials: &testCredentials,
	}

	rs, err = s.Identity.UpdateIdentity(req2, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

	res, err = s.Identity.QueryIdentity(id)
	require.NoError(s.T(), err)
}
