package nft

import (
	"errors"
	"strings"

	sdk "github.com/irisnet/core-sdk-go/types"
)

const (
	ModuleName = "nft"
)

var (
	_ sdk.Msg = &MsgIssueDenom{}
	_ sdk.Msg = &MsgTransferNFT{}
	_ sdk.Msg = &MsgEditNFT{}
	_ sdk.Msg = &MsgMintNFT{}
	_ sdk.Msg = &MsgBurnNFT{}
)

func (m MsgIssueDenom) ValidateBasic() error {
	if len(m.Sender) == 0 {
		return errors.New("missing sender address")
	}

	if err := sdk.ValidateAccAddress(m.Sender); err != nil {
		return err
	}
	id := strings.TrimSpace(m.Id)
	if len(id) == 0 {
		return errors.New("missing id")
	}
	return nil
}

func (m MsgIssueDenom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

func (m MsgTransferNFT) ValidateBasic() error {
	if len(m.Sender) == 0 {
		return errors.New("missing sender address")
	}
	if err := sdk.ValidateAccAddress(m.Sender); err != nil {
		return err
	}

	if len(m.Recipient) == 0 {
		return errors.New("missing recipient address")
	}
	if err := sdk.ValidateAccAddress(m.Recipient); err != nil {
		return err
	}

	denom := strings.TrimSpace(m.DenomId)
	if len(denom) == 0 {
		return errors.New("missing denom")
	}

	tokenID := strings.TrimSpace(m.Id)
	if len(tokenID) == 0 {
		return errors.New("missing ID")
	}
	return nil
}

func (m MsgTransferNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

func (m MsgEditNFT) ValidateBasic() error {
	if len(m.Sender) == 0 {
		return errors.New("missing sender address")
	}
	if err := sdk.ValidateAccAddress(m.Sender); err != nil {
		return err
	}

	denom := strings.TrimSpace(m.DenomId)
	if len(denom) == 0 {
		return errors.New("missing denom")
	}

	tokenID := strings.TrimSpace(m.Id)
	if len(tokenID) == 0 {
		return errors.New("missing ID")
	}
	return nil
}

func (m MsgEditNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

func (m MsgMintNFT) ValidateBasic() error {
	if len(m.Sender) == 0 {
		return errors.New("missing sender address")
	}
	if err := sdk.ValidateAccAddress(m.Sender); err != nil {
		return err
	}

	denom := strings.TrimSpace(m.DenomId)
	if len(denom) == 0 {
		return errors.New("missing denom")
	}

	tokenID := strings.TrimSpace(m.Id)
	if len(tokenID) == 0 {
		return errors.New("missing ID")
	}
	return nil
}

func (m MsgMintNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

func (m MsgBurnNFT) ValidateBasic() error {
	if len(m.Sender) == 0 {
		return errors.New("missing sender address")
	}
	if err := sdk.ValidateAccAddress(m.Sender); err != nil {
		return err
	}

	denom := strings.TrimSpace(m.DenomId)
	if len(denom) == 0 {
		return errors.New("missing denom")
	}

	tokenID := strings.TrimSpace(m.Id)
	if len(tokenID) == 0 {
		return errors.New("missing ID")
	}
	return nil
}

func (m MsgBurnNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

func (o Owner) Convert() interface{} {
	var idcs []IDC
	for _, idc := range o.IDCollections {
		idcs = append(idcs, IDC{
			Denom:    idc.DenomId,
			TokenIDs: idc.TokenIds,
		})
	}
	return QueryOwnerResp{
		Address: o.Address,
		IDCs:    idcs,
	}
}

func (this BaseNFT) Convert() interface{} {
	return QueryNFTResp{
		ID:      this.Id,
		Name:    this.Name,
		URI:     this.URI,
		Data:    this.Data,
		Creator: this.Owner,
	}
}

type NFTs []BaseNFT

func (this Denom) Convert() interface{} {
	return QueryDenomResp{
		ID:      this.Id,
		Name:    this.Name,
		Schema:  this.Schema,
		Creator: this.Creator,
	}
}

type denoms []Denom

func (this denoms) Convert() interface{} {
	var denoms []QueryDenomResp
	for _, denom := range this {
		denoms = append(denoms, denom.Convert().(QueryDenomResp))
	}
	return denoms
}

func (c Collection) Convert() interface{} {
	var nfts []QueryNFTResp
	for _, nft := range c.NFTs {
		nfts = append(nfts, QueryNFTResp{
			ID:      nft.Id,
			Name:    nft.Name,
			URI:     nft.URI,
			Data:    nft.Data,
			Creator: nft.Owner,
		})
	}
	return QueryCollectionResp{
		Denom: c.Denom.Convert().(QueryDenomResp),
		NFTs:  nfts,
	}
}
