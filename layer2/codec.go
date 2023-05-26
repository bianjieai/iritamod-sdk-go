package layer2

import (
	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	cryptocodec "github.com/irisnet/core-sdk-go/common/crypto/codec"
	sdk "github.com/irisnet/core-sdk-go/types"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateL2Space{},
		&MsgTransferL2Space{},
		&MsgCreateNFTs{},
		&MsgUpdateNFTs{},
		&MsgDeleteNFTs{},
		&MsgDepositClassForNFT{},
		&MsgWithdrawClassForNFT{},
		&MsgDepositTokenForNFT{},
		&MsgWithdrawTokenForNFT{},
	)
}