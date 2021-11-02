package node

import (
	"errors"
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

func (m MsgCreateValidator) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return errors.New("operator missing")
	}
	if len(strings.TrimSpace(m.Name)) == 0 {
		return errors.New("validator name cannot be blank")
	}

	if len(m.Certificate) == 0 {
		return errors.New("certificate missing")
	}
	if m.Power <= 0 {
		return errors.New("power must be positive")
	}
	return nil
}

func (m MsgCreateValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

func (m MsgUpdateValidator) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return errors.New("operator missing")
	}
	if len(m.Id) == 0 {
		return errors.New("validator id cannot be blank")
	}

	if m.Power < 0 {
		return errors.New("power can not be negative")
	}
	return nil
}

func (m MsgUpdateValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

func (m MsgRemoveValidator) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return errors.New("operator missing")
	}
	if len(m.Id) == 0 {
		return errors.New("validator id cannot be blank")
	}
	return nil
}

func (m MsgRemoveValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

func (m MsgGrantNode) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return errors.New("operator missing")
	}
	return nil
}

func (m MsgGrantNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

func (m MsgRevokeNode) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return errors.New("operator missing")
	}
	if len(m.Id) == 0 {
		return errors.New("validator id cannot be blank")
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
