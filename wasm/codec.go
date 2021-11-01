package wasm

import (
	"github.com/irisnet/core-sdk-go/codec"
	"github.com/irisnet/core-sdk-go/codec/types"
	cryptocodec "github.com/irisnet/core-sdk-go/crypto/codec"
	sdk "github.com/irisnet/core-sdk-go/types"
)

const (
	// ModuleName define the module name
	ModuleName = "wasm"
)

var (
	amino = codec.NewLegacyAmino()
	// ModuleCdc define the codec for wasm module
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

// RegisterInterfaces regisger the implement of the msg interface for InterfaceRegistry
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgStoreCode{},
		&MsgInstantiateContract{},
		&MsgExecuteContract{},
		&MsgMigrateContract{},
		&MsgUpdateAdmin{},
		&MsgClearAdmin{},
	)
}
