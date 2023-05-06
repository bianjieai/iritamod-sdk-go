package identity

import (
	"encoding/json"

	sdk "github.com/irisnet/core-sdk-go/types"
)

// Identity message types and params
const (
	TypeMsgCreateIdentity = "create_identity" // type for MsgCreateIdentity
	TypeMsgUpdateIdentity = "update_identity" // type for MsgUpdateIdentity

	IDLength     = 16  // size of the ID in bytes
	MaxURILength = 140 // maximum size of the URI

	DoNotModifyDesc = "[do-not-modify]" // description used to indicate not to modify a field

	ModuleName = "identity"
)

var (
	_ sdk.Msg = &MsgCreateIdentity{}
	_ sdk.Msg = &MsgUpdateIdentity{}
)

func (m MsgCreateIdentity) Route() string {
	return ModuleName
}

func (m MsgCreateIdentity) Type() string {
	return TypeMsgCreateIdentity
}

func (m MsgCreateIdentity) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements Msg
func (m MsgCreateIdentity) ValidateBasic() error {
	return ValidateIdentityFields(
		m.Id,
		m.PubKey,
		m.Certificate,
		m.Credentials,
		m.Owner,
		m.Data,
	)
}

// GetSigners implements Msg
func (m MsgCreateIdentity) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Owner)}
}

func (m MsgUpdateIdentity) Route() string {
	return ModuleName
}

func (m MsgUpdateIdentity) Type() string {
	return TypeMsgUpdateIdentity
}

func (m MsgUpdateIdentity) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements m.
func (m MsgUpdateIdentity) ValidateBasic() error {
	return ValidateIdentityFields(
		m.Id,
		m.PubKey,
		m.Certificate,
		m.Credentials,
		m.Owner,
		m.Data,
	)
}

// GetSigners implements m.
func (m MsgUpdateIdentity) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Owner)}
}

// ValidateIdentityFields validates the given identity fields
func ValidateIdentityFields(
	id string,
	pubKey *PubKeyInfo,
	certificate,
	credentials string,
	owner string,
	data string,
) error {
	if len(owner) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "owner missing")
	}

	if len(id) != IDLength*2 {
		return sdk.WrapWithMessage(ErrValidateBasic, "size of the ID must be %d in bytes", IDLength)
	}

	if len(credentials) > MaxURILength {
		return sdk.WrapWithMessage(ErrValidateBasic, "length of the credentials uri must not be greater than %d", MaxURILength)
	}

	return nil
}

func (m Identity) Convert() interface{} {
	var pubKeyInfos []PubKeyInfo
	for _, info := range m.PubKeys {
		pubKeyInfos = append(pubKeyInfos, PubKeyInfo{
			PubKey:    info.PubKey,
			Algorithm: info.Algorithm,
		})
	}

	return QueryIdentityResp{
		Id:           m.Id,
		PubKeyInfos:  pubKeyInfos,
		Certificates: m.Certificates,
		Credentials:  m.Credentials,
		Owner:        m.Owner,
	}
}

// MarshalJSON returns the JSON representation
func (p PubKeyAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(PubKeyAlgorithm_name[int32(p)])
}

// UnmarshalJSON unmarshals raw JSON bytes into a PubKeyAlgorithm
func (p *PubKeyAlgorithm) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return sdk.WrapWithMessage(ErrUnmarshal, err.Error())
	}

	algo := PubKeyAlgorithm_value[s]
	*p = PubKeyAlgorithm(algo)
	return nil
}
