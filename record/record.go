package record

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/bianjieai/iritamod-sdk-go/types"

	"github.com/irisnet/core-sdk-go/codec"
	codectypes "github.com/irisnet/core-sdk-go/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type recordClient struct {
	sdk.BaseClient
	codec.Codec
}

func NewClient(bc sdk.BaseClient, cdc codec.Codec) Client {
	return recordClient{
		BaseClient: bc,
		Codec:      cdc,
	}
}

func (r recordClient) Name() string {
	return ModuleName
}

func (r recordClient) RegisterInterfaceTypes(registry codectypes.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (r recordClient) CreateRecord(request CreateRecordRequest, baseTx sdk.BaseTx) (string, error) {
	creator, err := r.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return "", err
	}

	msg := &MsgCreateRecord{
		Contents: request.Contents,
		Creator:  creator.String(),
	}

	res, err := r.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return "", err
	}

	recordID, er := sdk.StringifyEvents(res.TxResult.Events).GetValue(eventTypeCreateRecord, attributeKeyRecordID)
	if er != nil {
		return "", er
	}

	return recordID, nil
}

func (r recordClient) QueryRecord(request QueryRecordReq) (QueryRecordResp, error) {
	rID, err := hex.DecodeString(request.RecordID)
	if err != nil {
		return QueryRecordResp{}, errors.New(fmt.Sprintf("invalid record id, must be hex encoded string,but got %s", request.RecordID))
	}

	recordKey := GetRecordKey(rID)

	res, err := r.QueryStore(recordKey, ModuleName, request.Height, request.Prove)
	if err != nil {
		return QueryRecordResp{}, err
	}

	var record Record
	if err := r.Codec.UnmarshalLengthPrefixed(res.Value, &record); err != nil {
		return QueryRecordResp{}, err
	}

	result := record.Convert().(QueryRecordResp)

	var proof []byte
	if request.Prove {
		proof = r.MustMarshalJSON(res.ProofOps)
	}

	result.Proof = types.ProofValue{
		Proof: proof,
		Path:  []string{ModuleName, string(recordKey)},
		Value: res.Value,
	}
	result.Height = res.Height
	return result, nil
}
