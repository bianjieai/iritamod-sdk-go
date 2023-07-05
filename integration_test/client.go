package integration

import (
	"github.com/bianjieai/iritamod-sdk-go/side-chain"
	"github.com/bianjieai/iritamod-sdk-go/slashing"
	"github.com/irisnet/core-sdk-go/bank"
	"github.com/irisnet/core-sdk-go/client"
	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/types"

	"github.com/bianjieai/iritamod-sdk-go/identity"
	"github.com/bianjieai/iritamod-sdk-go/node"
	"github.com/bianjieai/iritamod-sdk-go/params"
	"github.com/bianjieai/iritamod-sdk-go/perm"
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
	Key  client.Client
	Bank bank.Client

	Identity identity.Client
	Node     node.Client
	Params   params.Client
	Perm     perm.Client
	Slashing slashing.Client

	SideChainClient side_chain.Client
}

func NewClient(cfg types.ClientConfig) Client {
	encodingConfig := makeEncodingConfig()

	// create an instance of baseClient

	baseClient := client.NewBaseClient(cfg, encodingConfig, nil)
	bankClient := bank.NewClient(baseClient, encodingConfig.Marshaler)
	keysClient := client.NewKeysClient(cfg, baseClient)

	identityClient := identity.NewClient(baseClient, encodingConfig.Marshaler)
	nodeClient := node.NewClient(baseClient, encodingConfig.Marshaler)
	paramsClient := params.NewClient(baseClient, encodingConfig.Marshaler)
	permClient := perm.NewClient(baseClient, encodingConfig.Marshaler)
	slashingClient := slashing.NewClient(baseClient, encodingConfig.Marshaler)

	sideChainClient := side_chain.NewClient(baseClient, encodingConfig.Marshaler)

	client := &Client{
		logger:          baseClient.Logger(),
		BaseClient:      baseClient,
		moduleManager:   make(map[string]types.Module),
		encodingConfig:  encodingConfig,
		Key:             keysClient,
		Bank:            bankClient,
		Identity:        identityClient,
		Node:            nodeClient,
		Params:          paramsClient,
		Perm:            permClient,
		Slashing:        slashingClient,
		SideChainClient: sideChainClient,
	}

	client.RegisterModule(
		bankClient,
		identityClient,
		nodeClient,
		paramsClient,
		permClient,
		slashingClient,
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
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	txCfg := txtypes.NewTxConfig(protoCodec, txtypes.DefaultSignModes)

	encodingConfig := types.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         protoCodec,
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
