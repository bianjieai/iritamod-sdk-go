package token

import (
	json2 "encoding/json"
	"strconv"

	sdk "github.com/irisnet/core-sdk-go/types"
)

const (
	ModuleName = "token"
)

var (
	_ sdk.Msg = &MsgIssueToken{}
	_ sdk.Msg = &MsgEditToken{}
	_ sdk.Msg = &MsgMintToken{}
	_ sdk.Msg = &MsgTransferTokenOwner{}
)

func (m MsgIssueToken) Route() string {
	return ModuleName
}

func (m MsgIssueToken) Type() string {
	return "issue_token"
}

func (m MsgIssueToken) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

// ValidateBasic Implements Msg.
func (msg MsgIssueToken) ValidateBasic() error {
	if len(msg.Owner) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "owner must be not empty")
	}

	if err := sdk.ValidateAccAddress(msg.Owner); err != nil {
		return sdk.WrapWithMessage(ErrValidateAccAddress, err.Error())
	}

	if len(msg.Symbol) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "symbol must be not empty")
	}

	if len(msg.Name) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "name must be not empty")
	}

	if len(msg.MinUnit) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "minUnit must be not empty")
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgIssueToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Owner)}
}

func (m MsgTransferTokenOwner) Route() string {
	return ModuleName
}

func (m MsgTransferTokenOwner) Type() string {
	return "transfer_token_owner"
}

func (m MsgTransferTokenOwner) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgTransferTokenOwner) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.SrcOwner)}
}

func (msg MsgTransferTokenOwner) ValidateBasic() error {
	if len(msg.SrcOwner) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "srcOwner must be not empty")
	}

	if err := sdk.ValidateAccAddress(msg.SrcOwner); err != nil {
		return sdk.WrapWithMessage(ErrValidateAccAddress, err.Error())
	}

	if len(msg.DstOwner) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "dstOwner must be not empty")
	}

	if err := sdk.ValidateAccAddress(msg.DstOwner); err != nil {
		return sdk.WrapWithMessage(ErrValidateAccAddress, err.Error())
	}

	if len(msg.Symbol) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "symbol must be not empty")
	}

	return nil
}

func (m MsgEditToken) Route() string {
	return ModuleName
}

func (m MsgEditToken) Type() string {
	return "edit_token"
}

func (m MsgEditToken) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements Msg
func (msg MsgEditToken) ValidateBasic() error {
	if len(msg.Owner) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "owner must be not empty")
	}

	if err := sdk.ValidateAccAddress(msg.Owner); err != nil {
		return sdk.WrapWithMessage(ErrValidateAccAddress, err.Error())
	}

	if len(msg.Symbol) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "symbol must be not empty")
	}
	return nil
}

// GetSigners implements Msg
func (msg MsgEditToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Owner)}
}

func (m MsgMintToken) Route() string {
	return ModuleName
}

func (m MsgMintToken) Type() string {
	return "mint_token"
}

func (m MsgMintToken) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgMintToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Owner)}
}

// ValidateBasic implements Msg
func (msg MsgMintToken) ValidateBasic() error {
	if len(msg.Owner) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "owner must be not empty")
	}

	if err := sdk.ValidateAccAddress(msg.Owner); err != nil {
		return sdk.WrapWithMessage(ErrValidateAccAddress, err.Error())
	}

	if len(msg.Symbol) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "symbol must be not empty")
	}
	return nil
}

type Bool string

func (b Bool) ToBool() bool {
	v := string(b)
	if len(v) == 0 {
		return false
	}
	result, _ := strconv.ParseBool(v)
	return result
}

func (b Bool) String() string {
	return string(b)
}

// Marshal needed for protobuf compatibility
func (b Bool) Marshal() ([]byte, error) {
	return []byte(b), nil
}

// Unmarshal needed for protobuf compatibility
func (b *Bool) Unmarshal(data []byte) error {
	*b = Bool(data[:])
	return nil
}

// Marshals to JSON using string
func (b Bool) MarshalJSON() ([]byte, error) {
	return json2.Marshal(b.String())
}

// Unmarshals from JSON assuming Bech32 encoding
func (b *Bool) UnmarshalJSON(data []byte) error {
	var s string
	err := json2.Unmarshal(data, &s)
	if err != nil {
		return sdk.WrapWithMessage(ErrUnmarshal, err.Error())
	}
	*b = Bool(s)
	return nil
}

// GetSymbol implements exported.TokenI
func (t Token) GetSymbol() string {
	return t.Symbol
}

// GetName implements exported.TokenI
func (t Token) GetName() string {
	return t.Name
}

// GetScale implements exported.TokenI
func (t Token) GetScale() uint32 {
	return t.Scale
}

// GetMinUnit implements exported.TokenI
func (t Token) GetMinUnit() string {
	return t.MinUnit
}

// GetInitialSupply implements exported.TokenI
func (t Token) GetInitialSupply() uint64 {
	return t.InitialSupply
}

// GetMaxSupply implements exported.TokenI
func (t Token) GetMaxSupply() uint64 {
	return t.MaxSupply
}

// GetMintable implements exported.TokenI
func (t Token) GetMintable() bool {
	return t.Mintable
}

// GetOwner implements exported.TokenI
func (t Token) GetOwner() sdk.AccAddress {
	return sdk.MustAccAddressFromBech32(t.Owner)
}

func (t Token) Convert() interface{} {
	return sdk.Token{
		Symbol:        t.Symbol,
		Name:          t.Name,
		Scale:         t.Scale,
		MinUnit:       t.MinUnit,
		InitialSupply: t.InitialSupply,
		MaxSupply:     t.MaxSupply,
		Mintable:      t.Mintable,
		Owner:         t.Owner,
	}
}

type Tokens []TokenInterface

func (ts Tokens) Convert() interface{} {
	var tokens sdk.Tokens
	for _, t := range ts {
		tokens = append(tokens, sdk.Token{
			Symbol:        t.GetSymbol(),
			Name:          t.GetName(),
			Scale:         t.GetScale(),
			MinUnit:       t.GetMinUnit(),
			InitialSupply: t.GetInitialSupply(),
			MaxSupply:     t.GetMaxSupply(),
			Mintable:      t.GetMintable(),
			Owner:         t.GetOwner().String(),
		})
	}
	return tokens
}

type TokenInterface interface {
	GetSymbol() string
	GetName() string
	GetScale() uint32
	GetMinUnit() string
	GetInitialSupply() uint64
	GetMaxSupply() uint64
	GetMintable() bool
	GetOwner() sdk.AccAddress
}

func (p Params) Convert() interface{} {
	return QueryParamsResp{
		TokenTaxRate:      p.TokenTaxRate.String(),
		IssueTokenBaseFee: p.IssueTokenBaseFee.String(),
		MintTokenFeeRatio: p.MintTokenFeeRatio.String(),
	}
}

func (t QueryFeesResponse) Convert() interface{} {
	return QueryFeesResp(t)
}
