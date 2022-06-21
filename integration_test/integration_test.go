package integration

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/irisnet/core-sdk-go/common/log"
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/store"
)

const (
	nodeURI  = "tcp://localhost:26657"
	grpcAddr = "localhost:9090"
	chainID  = "iritasdk"
	charset  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	addr     = "iaa1d3p96vyvekh8wwrse5s2y7vhq9n6967mefsrav"
	name     = "node0"
	password = "iritasdk"
)

var (
	gasWanted = uint64(1000000)
	feeWanted = sdk.NewDecCoins(
		sdk.NewDecCoin(
			"ugas",
			sdk.NewInt(2000000)))
)

type IntegrationTestSuite struct {
	suite.Suite
	Client
	r            *rand.Rand
	rootAccount  MockAccount
	randAccounts []MockAccount
}

type SubTest struct {
	testName string
	testCase func(s IntegrationTestSuite)
}

// MockAccount define an account for test
type MockAccount struct {
	Name     string
	Password string
	Address  sdk.AccAddress
}

func TestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	feeCoin, err := sdk.ParseDecCoins("10000ugas")

	options := []sdk.Option{
		sdk.FeeOption(feeCoin),
		sdk.AlgoOption("sm2"),
		sdk.KeyDAOOption(store.NewMemory(nil)),
		sdk.TimeoutOption(6),
		sdk.CachedOption(true),
	}

	cfg, err := sdk.NewClientConfig(nodeURI, grpcAddr, chainID, options...)
	if err != nil {
		panic(err)
	}

	s.Client = NewClient(cfg)
	s.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	s.rootAccount = MockAccount{
		Name:     name,
		Password: password,
		Address:  sdk.MustAccAddressFromBech32(addr),
	}

	s.SetLogger(log.NewLogger(log.Config{
		Format: log.FormatText,
		Level:  log.DebugLevel,
	}))
	s.initAccount()
}

func (s *IntegrationTestSuite) initAccount() {

	address, err := s.Import(
		s.Account().Name,
		s.Account().Password,
		string(getPrivKeyArmor()))
	if err != nil {
		panic(err)
	}
	require.Equal(s.T(), address, addr)

	//var receipts bank.Receipts
	for i := 0; i < 5; i++ {
		name := s.RandStringOfLength(10)
		pwd := s.RandStringOfLength(16)
		address, _, err := s.Add(name, "1234567890")
		if err != nil {
			panic("generate test account fail")
		}

		s.randAccounts = append(s.randAccounts, MockAccount{
			Name:     name,
			Password: pwd,
			Address:  sdk.MustAccAddressFromBech32(address),
		})
	}
}

// RandStringOfLength return a random string
func (s *IntegrationTestSuite) RandStringOfLength(l int) string {
	var result []byte
	bytes := []byte(charset)
	for i := 0; i < l; i++ {
		result = append(result, bytes[s.r.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandAccount return a random test account
func (s *IntegrationTestSuite) GetRandAccount() MockAccount {
	return s.randAccounts[s.r.Intn(len(s.randAccounts))]
}

// Account return a test account
func (s *IntegrationTestSuite) Account() MockAccount {
	return s.rootAccount
}

func getPrivKeyArmor() []byte {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = filepath.Dir(path)
	path = filepath.Join(path, "integration_test/scripts/priv.key")
	bz, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bz
}
