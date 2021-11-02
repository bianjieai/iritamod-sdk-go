// Package token allows individuals and companies to create and issue their own tokens.
//

package token

import (
	"context"
	"strconv"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/irisnet/core-sdk-go/codec"
	"github.com/irisnet/core-sdk-go/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/errors"
	"github.com/irisnet/core-sdk-go/types/query"
)

type tokenClient struct {
	sdk.BaseClient
	codec.Codec
}

func NewClient(bc sdk.BaseClient, cdc codec.Codec) Client {
	return tokenClient{
		BaseClient: bc,
		Codec:      cdc,
	}
}

func (t tokenClient) Name() string {
	return ModuleName
}

func (t tokenClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (t tokenClient) IssueToken(req IssueTokenRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	owner, err := t.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrQueryAddress, err.Error())
	}

	msg := &MsgIssueToken{
		Symbol:        req.Symbol,
		Name:          req.Name,
		Scale:         req.Scale,
		MinUnit:       req.MinUnit,
		InitialSupply: req.InitialSupply,
		MaxSupply:     req.MaxSupply,
		Mintable:      req.Mintable,
		Owner:         owner.String(),
	}
	send, err := t.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (t tokenClient) EditToken(req EditTokenRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	owner, err := t.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrQueryAddress, err.Error())
	}

	msg := &MsgEditToken{
		Symbol:    req.Symbol,
		Name:      req.Name,
		MaxSupply: req.MaxSupply,
		Mintable:  Bool(strconv.FormatBool(req.Mintable)),
		Owner:     owner.String(),
	}

	send, err := t.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (t tokenClient) TransferToken(to string, symbol string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	owner, err := t.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrQueryAddress, err.Error())
	}

	if err := sdk.ValidateAccAddress(to); err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrValidateAccAddress, err.Error())
	}

	msg := &MsgTransferTokenOwner{
		SrcOwner: owner.String(),
		DstOwner: to,
		Symbol:   symbol,
	}
	send, err := t.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (t tokenClient) MintToken(symbol string, amount uint64, to string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	owner, err := t.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrQueryAddress, err.Error())
	}

	receipt := owner.String()
	if len(to) != 0 {
		if err := sdk.ValidateAccAddress(to); err != nil {
			return ctypes.ResultTx{}, errors.Wrap(ErrValidateAccAddress, err.Error())
		} else {
			receipt = to
		}
	}

	msg := &MsgMintToken{
		Symbol: symbol,
		Amount: amount,
		To:     receipt,
		Owner:  owner.String(),
	}
	send, err := t.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, errors.Wrap(ErrBuildAndSend, err.Error())
	}
	return send, nil
}

func (t tokenClient) QueryToken(denom string) (sdk.Token, error) {
	conn, err := t.GenConn()
	if err != nil {
		return sdk.Token{}, errors.Wrap(ErrGenConn, err.Error())
	}

	request := &QueryTokenRequest{
		Denom: denom,
	}
	res, err := NewQueryClient(conn).Token(context.Background(), request)
	if err != nil {
		return sdk.Token{}, errors.Wrap(ErrQueryToken, err.Error())
	}
	var evi TokenInterface
	if err = t.UnpackAny(res.Token, &evi); err != nil {
		return sdk.Token{}, errors.Wrap(errors.ErrUnpackAny, err.Error())
	}
	tokens := make(Tokens, 0)
	tokens = append(tokens, evi.(*Token))
	ts := tokens.Convert().(sdk.Tokens)
	t.SaveTokens(ts...)
	return ts[0], nil
}

func (t tokenClient) QueryTokens(owner string, pageReq *query.PageRequest) (sdk.Tokens, error) {
	var ownerAddr string
	if len(owner) > 0 {
		if err := sdk.ValidateAccAddress(owner); err != nil {
			return nil, errors.Wrap(ErrValidateAccAddress, err.Error())
		}
		ownerAddr = owner
	}

	conn, err := t.GenConn()

	if err != nil {
		return sdk.Tokens{}, errors.Wrap(ErrGenConn, err.Error())
	}

	request := &QueryTokensRequest{
		Owner:      ownerAddr,
		Pagination: pageReq,
	}

	res, err := NewQueryClient(conn).Tokens(context.Background(), request)
	if err != nil {
		return sdk.Tokens{}, errors.Wrap(ErrQueryToken, err.Error())
	}

	tokens := make(Tokens, 0, len(res.Tokens))
	for _, eviAny := range res.Tokens {
		var evi TokenInterface
		if err = t.UnpackAny(eviAny, &evi); err != nil {
			return sdk.Tokens{}, errors.Wrap(errors.ErrUnpackAny, err.Error())
		}
		tokens = append(tokens, evi.(*Token))
	}

	ts := tokens.Convert().(sdk.Tokens)
	t.SaveTokens(ts...)
	return ts, nil
}

func (t tokenClient) QueryFees(symbol string) (QueryFeesResp, error) {
	conn, err := t.GenConn()

	if err != nil {
		return QueryFeesResp{}, errors.Wrap(ErrGenConn, err.Error())
	}

	request := &QueryFeesRequest{
		Symbol: symbol,
	}

	res, err := NewQueryClient(conn).Fees(context.Background(), request)
	if err != nil {
		return QueryFeesResp{}, errors.Wrap(ErrQueryToken, err.Error())
	}

	return res.Convert().(QueryFeesResp), nil
}

func (t tokenClient) QueryParams() (QueryParamsResp, error) {
	conn, err := t.GenConn()

	if err != nil {
		return QueryParamsResp{}, errors.Wrap(ErrGenConn, err.Error())
	}

	res, err := NewQueryClient(conn).Params(
		context.Background(),
		&QueryParamsRequest{},
	)
	if err != nil {
		return QueryParamsResp{}, errors.Wrap(ErrQueryToken, err.Error())
	}

	return res.Params.Convert().(QueryParamsResp), nil
}
