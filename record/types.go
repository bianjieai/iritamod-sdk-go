package record

import (
	"github.com/irisnet/core-sdk-go/types/errors"

	sdk "github.com/irisnet/core-sdk-go/types"
)

const (
	ModuleName = "record"

	attributeKeyRecordID  = "record_id"
	eventTypeCreateRecord = "create_record"
)

var (
	_ sdk.Msg = &MsgCreateRecord{}

	recordKey = []byte{0x01} // record key
)

// ValidateBasic implements Msg.
func (msg MsgCreateRecord) ValidateBasic() error {
	if len(msg.Contents) == 0 {
		return errors.Wrap(ErrValidateBasic, "contents missing")
	}
	if len(msg.Creator) == 0 {
		return errors.Wrap(ErrValidateBasic, "creator missing")
	}

	if err := sdk.ValidateAccAddress(msg.Creator); err != nil {
		return err
	}

	for i, content := range msg.Contents {
		if len(content.Digest) == 0 {
			return errors.Wrapf(ErrValidateBasic, "content[%d] digest missing", i)
		}
		if len(content.DigestAlgo) == 0 {
			return errors.Wrapf(ErrValidateBasic, "content[%d] digest algo missing", i)
		}
	}
	return nil
}

// GetSigners implements Msg.
func (msg MsgCreateRecord) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (this Record) Convert() interface{} {
	return QueryRecordResp{
		Record: Data{
			TxHash:   this.TxHash,
			Contents: this.Contents,
			Creator:  this.Creator,
		},
	}
}

// GetRecordKey returns record key bytes
func GetRecordKey(recordID []byte) []byte {
	return append(recordKey, recordID...)
}
