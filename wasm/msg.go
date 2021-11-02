package wasm

import (
	"encoding/json"
	"errors"

	sdk "github.com/irisnet/core-sdk-go/types"
)

// message types for the wasm client
const (
	RouterKey                  string = "wasm"
	TypeMsgStoreCode           string = "store-code"
	TypeMsgInstantiateContract string = "instantiate"
	TypeMsgExecuteContract     string = "execute"
	TypeMsgMigrateContract     string = "migrate"
	TypeUpdateAdmin            string = "update-contract-admin"
	TypeClearAdmin             string = "clear-contract-admin"
)

var (
	_ sdk.Msg = &MsgStoreCode{}
	_ sdk.Msg = &MsgInstantiateContract{}
	_ sdk.Msg = &MsgExecuteContract{}
	_ sdk.Msg = &MsgMigrateContract{}
	_ sdk.Msg = &MsgUpdateAdmin{}
	_ sdk.Msg = &MsgClearAdmin{}
)

// ValidateBasic implement sdk.Msg
func (msg MsgStoreCode) ValidateBasic() error {
	if err := sdk.ValidateAccAddress(msg.Sender); err != nil {
		return err
	}

	if len(msg.WASMByteCode) == 0 {
		return errors.New("WASMByteCode should not be empty")
	}

	return nil
}

// GetSigners implement sdk.Msg
func (msg MsgStoreCode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Sender)}
}

// ValidateBasic implement sdk.Msg
func (msg MsgInstantiateContract) ValidateBasic() error {
	if msg.CodeID == 0 {
		return errors.New("code id is required")
	}
	if msg.Label == "" {
		return errors.New("label is required")
	}
	if len(msg.Admin) != 0 {
		if err := sdk.ValidateAccAddress(msg.Admin); err != nil {
			return err
		}
	}
	if !json.Valid(msg.InitMsg) {
		return errors.New("InitMsg is not valid json")
	}
	return sdk.ValidateAccAddress(msg.Sender)
}

// GetSigners implement sdk.Msg
func (msg MsgInstantiateContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Sender)}
}

// ValidateBasic implement sdk.Msg
func (msg MsgExecuteContract) ValidateBasic() error {
	if err := sdk.ValidateAccAddress(msg.Contract); err != nil {
		return err
	}
	if !json.Valid(msg.Msg) {
		return errors.New("InitMsg is not valid json")
	}
	return sdk.ValidateAccAddress(msg.Sender)
}

// GetSigners implement sdk.Msg
func (msg MsgExecuteContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Sender)}
}

// ValidateBasic implement sdk.Msg
func (msg MsgMigrateContract) ValidateBasic() error {
	if msg.CodeID == 0 {
		return errors.New("code id is required")
	}

	if err := sdk.ValidateAccAddress(msg.Contract); err != nil {
		return err
	}

	if !json.Valid(msg.MigrateMsg) {
		return errors.New("migrate msg json")
	}
	return sdk.ValidateAccAddress(msg.Sender)
}

// GetSigners implement sdk.Msg
func (msg MsgMigrateContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Sender)}
}

// ValidateBasic implement sdk.Msg
func (msg MsgUpdateAdmin) ValidateBasic() error {
	return sdk.ValidateAccAddress(msg.Sender)
}

// GetSigners implement sdk.Msg
func (msg MsgUpdateAdmin) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Sender)}
}

// ValidateBasic implement sdk.Msg
func (msg MsgClearAdmin) ValidateBasic() error {
	return sdk.ValidateAccAddress(msg.Sender)
}

// GetSigners implement sdk.Msg
func (msg MsgClearAdmin) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Sender)}
}
