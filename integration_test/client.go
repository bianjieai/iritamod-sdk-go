package integration

import (
	"github.com/bianjieai/iritamod-sdk-go/identity"
	"github.com/bianjieai/iritamod-sdk-go/nft"
	"github.com/bianjieai/iritamod-sdk-go/node"
	"github.com/bianjieai/iritamod-sdk-go/oracle"
	"github.com/bianjieai/iritamod-sdk-go/params"
	"github.com/bianjieai/iritamod-sdk-go/perm"
	"github.com/bianjieai/iritamod-sdk-go/record"
	"github.com/bianjieai/iritamod-sdk-go/service"
	"github.com/bianjieai/iritamod-sdk-go/token"
	"github.com/bianjieai/iritamod-sdk-go/wasm"
	"github.com/irisnet/core-sdk-go/bank"
	"github.com/irisnet/core-sdk-go/client"
	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/gov"
	"github.com/irisnet/core-sdk-go/staking"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/tendermint/tendermint/libs/log"

	cdctypes "github.com/irisnet/core-sdk-go/common/codec/types"
	cryptocodec "github.com/irisnet/core-sdk-go/common/crypto/codec"
	txtypes "github.com/irisnet/core-sdk-go/types/tx"
)

type Client struct {
	logger         log.Logger
	moduleManager  map[string]types.Module
	encodingConfig types.EncodingConfig

	types.BaseClient
	Bank     bank.Client
	Staking  staking.Client
	Gov      gov.Client
	Identity identity.Client
	Nft      nft.Client
	Node     node.Client
	Service  service.Client
	Oracle   oracle.Client
	Params   params.Client
	Perm     perm.Client
	Record   record.Client
	Token    token.Client
	Wasm     wasm.Client
}

func NewClient(cfg types.ClientConfig) Client {
	encodingConfig := makeEncodingConfig()

	// create a instance of baseClient

	baseClient := client.NewBaseClient(cfg, encodingConfig, nil)
	bankClient := bank.NewClient(baseClient, encodingConfig.Marshaler)

	stakingClient := staking.NewClient(baseClient, encodingConfig.Marshaler)
	govClient := gov.NewClient(baseClient, encodingConfig.Marshaler)
	identityClient := identity.NewClient(baseClient, encodingConfig.Marshaler)
	nftClient := nft.NewClient(baseClient, encodingConfig.Marshaler)
	nodeClient := node.NewClient(baseClient, encodingConfig.Marshaler)
	serviceClient := service.NewClient(baseClient, encodingConfig.Marshaler)
	oracleClient := oracle.NewClient(baseClient, encodingConfig.Marshaler)
	paramsClient := params.NewClient(baseClient, encodingConfig.Marshaler)
	permClient := perm.NewClient(baseClient, encodingConfig.Marshaler)
	recordClient := record.NewClient(baseClient, encodingConfig.Marshaler)
	tokenClient := token.NewClient(baseClient, encodingConfig.Marshaler)
	wasmClient := wasm.NewClient(baseClient)

	client := &Client{
		logger:         baseClient.Logger(),
		BaseClient:     baseClient,
		moduleManager:  make(map[string]types.Module),
		encodingConfig: encodingConfig,
		Bank:           bankClient,
		Staking:        stakingClient,
		Gov:            govClient,
		Identity:       identityClient,
		Nft:            nftClient,
		Node:           nodeClient,
		Service:        serviceClient,
		Oracle:         oracleClient,
		Params:         paramsClient,
		Perm:           permClient,
		Record:         recordClient,
		Token:          tokenClient,
		Wasm:           wasmClient,
	}

	client.RegisterModule(
		bankClient,
		stakingClient,
		govClient,
		identityClient,
		nftClient,
		nodeClient,
		serviceClient,
		oracleClient,
		paramsClient,
		permClient,
		recordClient,
		tokenClient,
		wasmClient,
	)
	return *client
}

func (client Client) SetLogger(logger log.Logger) {
	client.BaseClient.SetLogger(logger)
}

func (client Client) Codec() *codec.LegacyAmino {
	return client.encodingConfig.Amino
}

func (client Client) AppCodec() codec.Marshaler {
	return client.encodingConfig.Marshaler
}

func (client Client) EncodingConfig() types.EncodingConfig {
	return client.encodingConfig
}

func (client Client) Manager() types.BaseClient {
	return client.BaseClient
}

func (client Client) RegisterModule(ms ...types.Module) {
	for _, m := range ms {
		m.RegisterInterfaceTypes(client.encodingConfig.InterfaceRegistry)
	}
}

func (client Client) Module(name string) types.Module {
	return client.moduleManager[name]
}

func makeEncodingConfig() types.EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := txtypes.NewTxConfig(marshaler, txtypes.DefaultSignModes)

	encodingConfig := types.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
	RegisterLegacyAminoCodec(encodingConfig.Amino)
	RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

// RegisterLegacyAminoCodec registers the sdk message type.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*types.Msg)(nil), nil)
	cdc.RegisterInterface((*types.Tx)(nil), nil)
	cryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers the sdk message type.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface("cosmos.v1beta1.Msg", (*types.Msg)(nil))
	txtypes.RegisterInterfaces(registry)
	cryptocodec.RegisterInterfaces(registry)
}
