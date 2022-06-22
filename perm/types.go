package perm

import (
	"encoding/json"
	"fmt"

	sdk "github.com/irisnet/core-sdk-go/types"
)

const (
	// ModuleName is the name of the perm module
	ModuleName             = "perm"
	TypeMsgAssignRoles     = "assign_roles"     // type for MsgAssignRoles
	TypeMsgUnassignRoles   = "unassign_roles"   // type for MsgUnassignRoles
	TypeMsgBlockAccount    = "block_account"    // type for MsgBlockAccount
	TypeMsgUnblockAccount  = "unblock_account"  // type for MsgUnblockAccount
	TypeMsgBlockContract   = "block_contract"   // type for MsgBlockContract
	TypeMsgUnblockContract = "unblock_contract" // type for MsgUnblockContract
)

var (
	_ sdk.Msg = &MsgAssignRoles{}
	_ sdk.Msg = &MsgUnassignRoles{}
	_ sdk.Msg = &MsgBlockAccount{}
	_ sdk.Msg = &MsgUnblockAccount{}
	_ sdk.Msg = &MsgBlockContract{}
	_ sdk.Msg = &MsgUnblockContract{}
)

func (m MsgAssignRoles) Route() string {
	return ModuleName
}

func (m MsgAssignRoles) Type() string {
	return TypeMsgAssignRoles
}

func (m MsgAssignRoles) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgAssignRoles) ValidateBasic() error {
	if len(m.Address) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "address missing")
	}

	if len(m.Operator) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "operator missing")
	}
	if len(m.Roles) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "roles missing")
	}
	return nil
}

func (m MsgAssignRoles) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

func (m MsgUnassignRoles) Route() string {
	return ModuleName
}

func (m MsgUnassignRoles) Type() string {
	return TypeMsgUnassignRoles
}

func (m MsgUnassignRoles) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgUnassignRoles) ValidateBasic() error {
	if len(m.Address) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "address missing")
	}
	if len(m.Operator) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "operator missing")
	}
	if len(m.Roles) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "roles missing")
	}
	return nil
}

func (m MsgUnassignRoles) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

func (m MsgBlockAccount) Route() string {
	return ModuleName
}

func (m MsgBlockAccount) Type() string {
	return TypeMsgBlockAccount
}

func (m MsgBlockAccount) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgBlockAccount) ValidateBasic() error {
	if len(m.Address) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "address missing")
	}
	if len(m.Operator) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "operator missing")
	}
	return nil
}

func (m MsgBlockAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

func (m MsgUnblockAccount) Route() string {
	return ModuleName
}

func (m MsgUnblockAccount) Type() string {
	return TypeMsgUnblockAccount
}

func (m MsgUnblockAccount) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgUnblockAccount) ValidateBasic() error {
	if len(m.Address) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "address missing")
	}
	if len(m.Operator) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "operator missing")
	}
	return nil
}

func (m MsgUnblockAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

// MsgBlockContract implements sdk.Msg

func (m MsgBlockContract) Route() string {
	return ModuleName
}

func (m MsgBlockContract) Type() string {
	return TypeMsgBlockContract
}

func (m MsgBlockContract) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgBlockContract) ValidateBasic() error {
	if len(m.ContractAddress) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "address missing")
	}
	if len(m.Operator) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "operator missing")
	}
	return nil
}

func (m MsgBlockContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

// MsgBlockContract implements sdk.Msg

func (m MsgUnblockContract) Route() string {
	return ModuleName
}

func (m MsgUnblockContract) Type() string {
	return TypeMsgUnblockContract
}

func (m MsgUnblockContract) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgUnblockContract) ValidateBasic() error {
	if len(m.ContractAddress) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "address missing")
	}
	if len(m.Operator) == 0 {
		return sdk.WrapWithMessage(ErrValidateBasic, "operator missing")
	}
	return nil
}

func (m MsgUnblockContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Operator)}
}

// RoleFromstring turn a string into an Auth
func roleFromString(str string) (Role, error) {
	switch str {
	case "RootAdmin":
		return RoleRootAdmin, nil

	case "PermAdmin":
		return RolePermAdmin, nil

	case "BlacklistAdmin":
		return RoleBlacklistAdmin, nil

	case "NodeAdmin":
		return RoleNodeAdmin, nil

	case "ParamAdmin":
		return RoleParamAdmin, nil

	case "PowerUser":
		return RolePowerUser, nil

	case "RelayerUser":
		return RoleRelayerUser, nil

	case "IDAdmin":
		return RoleIDAdmin, nil

	case "BaseM1Admin":
		return RoleBaseM1Admin, nil

	case "PlatformUser":
		return RolePlatformUser, nil

	default:
		return Role(0xff), fmt.Errorf("'%s' is not a valid role", str)
	}
}

// Marshal needed for protobuf compatibility
func (r Role) Marshal() ([]byte, error) {
	return []byte{byte(r)}, nil
}

// Unmarshal needed for protobuf compatibility
func (r *Role) Unmarshal(data []byte) error {
	*r = Role(data[0])
	return nil
}

// MarshalJSON Marshals to JSON using string representation of the status
func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.string())
}

// UnmarshalJSON Unmarshals from JSON assuming Bech32 encoding
func (r *Role) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	bz2, err := roleFromString(s)
	if err != nil {
		return err
	}

	*r = bz2
	return nil
}

// string implements the stringer interface.
func (r Role) string() string {
	switch r {
	case RoleRootAdmin:
		return "RootAdmin"

	case RolePermAdmin:
		return "PermAdmin"

	case RoleBlacklistAdmin:
		return "BlacklistAdmin"

	case RoleNodeAdmin:
		return "NodeAdmin"

	case RoleParamAdmin:
		return "ParamAdmin"

	case RolePowerUser:
		return "PowerUser"

	case RoleRelayerUser:
		return "RelayerUser"

	case RoleIDAdmin:
		return "IDAdmin"

	case RoleBaseM1Admin:
		return "BaseM1Admin"

	case RolePlatformUser:
		return "PlatformUser"

	default:
		return ""
	}
}
