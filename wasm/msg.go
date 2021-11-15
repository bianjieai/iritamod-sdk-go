package wasm

import (
	"fmt"

	"github.com/irisnet/core-sdk-go/types/errors"

	sdk "github.com/irisnet/core-sdk-go/types"
)

// message types for the wasm client
const (
	RouterKey                  string = "wasm"
	TypeMsgStoreCode           string = "store_code"
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
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return err
	}

	if err := validateWasmCode(msg.WASMByteCode); err != nil {
		return errors.Wrapf(errors.ErrInvalidRequest, "code bytes %s", err.Error())
	}

	if msg.InstantiatePermission != nil {
		if err := msg.InstantiatePermission.ValidateBasic(); err != nil {
			return errors.Wrap(err, "instantiate permission")
		}
	}
	return nil
}

// GetSigners implement sdk.Msg
func (msg MsgStoreCode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Sender)}
}

// ValidateBasic implement sdk.Msg
func (msg MsgInstantiateContract) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.Wrap(err, "sender")
	}

	if msg.CodeID == 0 {
		return errors.Wrap(errors.ErrInvalidRequest, "code id is required")
	}

	if err := validateLabel(msg.Label); err != nil {
		return errors.Wrap(errors.ErrInvalidRequest, "label is required")

	}

	if !msg.Funds.IsValid() {
		return ErrInvalidCoins
	}

	if len(msg.Admin) != 0 {
		if _, err := sdk.AccAddressFromBech32(msg.Admin); err != nil {
			return errors.Wrap(err, "admin")
		}
	}
	if err := msg.Msg.ValidateBasic(); err != nil {
		return errors.Wrap(err, "payload msg")
	}
	return nil
}

// GetSigners implement sdk.Msg
func (msg MsgInstantiateContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Sender)}
}

// ValidateBasic implement sdk.Msg
func (msg MsgExecuteContract) ValidateBasic() error {

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.Wrap(err, "sender")
	}
	if _, err := sdk.AccAddressFromBech32(msg.Contract); err != nil {
		return errors.Wrap(err, "contract")
	}

	if !msg.Funds.IsValid() {
		return errors.Wrap(ErrInvalidCoins, "sentFunds")
	}
	if err := msg.Msg.ValidateBasic(); err != nil {
		return errors.Wrap(err, "payload msg")
	}
	return nil
}

// GetSigners implement sdk.Msg
func (msg MsgExecuteContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Sender)}
}

// ValidateBasic implement sdk.Msg
func (msg MsgMigrateContract) ValidateBasic() error {
	if msg.CodeID == 0 {
		return errors.Wrap(errors.ErrInvalidRequest, "code id is required")
	}
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.Wrap(err, "sender")
	}
	if _, err := sdk.AccAddressFromBech32(msg.Contract); err != nil {
		return errors.Wrap(err, "contract")
	}

	if err := msg.Msg.ValidateBasic(); err != nil {
		return errors.Wrap(err, "payload msg")
	}

	return nil
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
	err := sdk.ValidateAccAddress(msg.Sender)
	if err != nil {
		return errors.Wrap(ErrValidateAccAddress, err.Error())
	}
	return nil
}

// GetSigners implement sdk.Msg
func (msg MsgClearAdmin) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Sender)}
}

// String implements the Stringer interface.
func (p MigrateContractProposal) String() string {
	return fmt.Sprintf(`Migrate Contract Proposal:
  Title:       %s
  Description: %s
  Contract:    %s
  Code id:     %d
  Run as:      %s
  Msg          %q
`, p.Title, p.Description, p.Contract, p.CodeID, p.RunAs, p.Msg)
}

// String implements the Stringer interface.
func (p StoreCodeProposal) String() string {
	return fmt.Sprintf(`Store Code Proposal:
  Title:       %s
  Description: %s
  Run as:      %s
  WasmCode:    %X
`, p.Title, p.Description, p.RunAs, p.WASMByteCode)
}

// String implements the Stringer interface.
func (p InstantiateContractProposal) String() string {
	return fmt.Sprintf(`Instantiate Code Proposal:
  Title:       %s
  Description: %s
  Run as:      %s
  Admin:       %s
  Code id:     %d
  Label:       %s
  Msg:         %q
  Funds:       %s
`, p.Title, p.Description, p.RunAs, p.Admin, p.CodeID, p.Label, p.Msg, p.Funds)
}

// String implements the Stringer interface.
func (p UpdateAdminProposal) String() string {
	return fmt.Sprintf(`Update Contract Admin Proposal:
  Title:       %s
  Description: %s
  Contract:    %s
  New Admin:   %s
`, p.Title, p.Description, p.Contract, p.NewAdmin)
}

// String implements the Stringer interface.
func (p ClearAdminProposal) String() string {
	return fmt.Sprintf(`Clear Contract Admin Proposal:
  Title:       %s
  Description: %s
  Contract:    %s
`, p.Title, p.Description, p.Contract)
}

// String implements the Stringer interface.
func (p PinCodesProposal) String() string {
	return fmt.Sprintf(`Pin Wasm Codes Proposal:
  Title:       %s
  Description: %s
  Codes:       %v
`, p.Title, p.Description, p.CodeIDs)
}

// String implements the Stringer interface.
func (p UnpinCodesProposal) String() string {
	return fmt.Sprintf(`Unpin Wasm Codes Proposal:
  Title:       %s
  Description: %s
  Codes:       %v
`, p.Title, p.Description, p.CodeIDs)
}
