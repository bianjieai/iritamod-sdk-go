package node

import (
	"strings"

	sdk "github.com/irisnet/core-sdk-go/types"
)

const (
	ModuleName = "node"
)

var (
	_ sdk.Msg = &MsgCreateValidator{}
	_ sdk.Msg = &MsgRemoveValidator{}
	_ sdk.Msg = &MsgUpdateValidator{}
	_ sdk.Msg = &MsgGrantNode{}
	_ sdk.Msg = &MsgRevokeNode{}
)

func (m MsgCreateValidator) Route() string {
	return ModuleName
}

func (m MsgCreateValidator) Type() string {
	return "create_validator"
}

func (m MsgCreateValidator) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgCreateValidator) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "operator missing")
	}
	if len(strings.TrimSpace(m.Name)) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "validator name cannot be blank")
	}

	if len(m.Certificate) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "certificate missing")
	}
	if m.Power <= 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "power must be positive")
	}
	return nil
}

func (m MsgCreateValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

func (m MsgUpdateValidator) Route() string {
	return ModuleName
}

func (m MsgUpdateValidator) Type() string {
	return "update_validator"
}

func (m MsgUpdateValidator) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgUpdateValidator) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "operator missing")
	}
	if len(m.Id) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "validator id cannot be blank")
	}

	if m.Power < 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "power can not be negative")
	}
	return nil
}

func (m MsgUpdateValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

func (m MsgRemoveValidator) Route() string {
	return ModuleName
}

func (m MsgRemoveValidator) Type() string {
	return "remove_validator"
}

func (m MsgRemoveValidator) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgRemoveValidator) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "operator missing")
	}
	if len(m.Id) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "validator id cannot be blank")
	}
	return nil
}

func (m MsgRemoveValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

func (m MsgGrantNode) Route() string {
	return ModuleName
}

func (m MsgGrantNode) Type() string {
	return "grant_node"
}

func (m MsgGrantNode) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgGrantNode) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "operator missing")
	}
	return nil
}

func (m MsgGrantNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

func (m MsgRevokeNode) Route() string {
	return ModuleName
}

func (m MsgRevokeNode) Type() string {
	return "revoke_node"
}

func (m MsgRevokeNode) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgRevokeNode) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "operator missing")
	}
	if len(m.Id) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "validator id cannot be blank")
	}
	return nil
}

func (m MsgRevokeNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

func (v Validator) Convert() interface{} {
	return QueryValidatorResp{
		ID:          v.Id,
		Name:        v.Name,
		Pubkey:      v.Pubkey,
		Certificate: v.Certificate,
		Power:       v.Power,
		Details:     v.Description,
		Jailed:      v.Jailed,
		Operator:    v.Operator,
	}
}

type validators []Validator

func (vs validators) Convert() interface{} {
	var vrs []QueryValidatorResp
	for _, v := range vs {
		vrs = append(vrs, QueryValidatorResp{
			ID:          v.Id,
			Name:        v.Name,
			Pubkey:      v.Pubkey,
			Certificate: v.Certificate,
			Power:       v.Power,
			Details:     v.Description,
			Jailed:      v.Jailed,
			Operator:    v.Operator,
		})
	}
	return vrs
}

func (n Node) Convert() interface{} {
	return QueryNodeResp{
		ID:          n.Id,
		Name:        n.Name,
		Certificate: n.Certificate,
	}
}

type nodes []Node

func (ns nodes) Convert() interface{} {
	var nrs []QueryNodeResp
	for _, n := range ns {
		nrs = append(nrs, QueryNodeResp{
			ID:          n.Id,
			Name:        n.Name,
			Certificate: n.Certificate,
		})
	}
	return nrs
}

func (p Params) Convert() interface{} {
	return QueryParamsResp(p)
}
