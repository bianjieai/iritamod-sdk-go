package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/irisnet/core-sdk-go/codec"
	"github.com/irisnet/core-sdk-go/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/query"
)

type serviceClient struct {
	sdk.BaseClient
	codec.Codec
}

func NewClient(bc sdk.BaseClient, cdc codec.Codec) Client {
	return serviceClient{
		BaseClient: bc,
		Codec:      cdc,
	}
}

func (s serviceClient) Name() string {
	return ModuleName
}

func (s serviceClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

//DefineService is responsible for creating a new service definition
func (s serviceClient) DefineService(request DefineServiceRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	author, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}
	msg := &MsgDefineService{
		Name:              request.ServiceName,
		Description:       request.Description,
		Tags:              request.Tags,
		Author:            author.String(),
		AuthorDescription: request.AuthorDescription,
		Schemas:           request.Schemas,
	}
	return s.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

//BindService is responsible for binding a new service definition
func (s serviceClient) BindService(request BindServiceRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	owner, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	var provider = owner.String()
	if len(request.Provider) > 0 {
		if err := sdk.ValidateAccAddress(request.Provider); err != nil {
			return ctypes.ResultTx{}, err
		}
		provider = request.Provider
	}

	amt, err := s.ToMinCoin(request.Deposit...)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	msg := &MsgBindService{
		ServiceName: request.ServiceName,
		Provider:    provider,
		Deposit:     amt,
		Pricing:     request.Pricing,
		QoS:         request.QoS,
		Options:     request.Options,
		Owner:       owner.String(),
	}
	return s.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

//UpdateServiceBinding updates the specified service binding
func (s serviceClient) UpdateServiceBinding(request UpdateServiceBindingRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	owner, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	var provider = owner.String()
	if len(request.Provider) > 0 {
		if err := sdk.ValidateAccAddress(request.Provider); err != nil {
			return ctypes.ResultTx{}, err
		}
		provider = request.Provider
	}

	amt, err := s.ToMinCoin(request.Deposit...)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	msg := &MsgUpdateServiceBinding{
		ServiceName: request.ServiceName,
		Provider:    provider,
		Deposit:     amt,
		Pricing:     request.Pricing,
		QoS:         request.QoS,
		Owner:       owner.String(),
	}
	return s.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

// DisableServiceBinding disables the specified service binding
func (s serviceClient) DisableServiceBinding(serviceName, provider string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	owner, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	var providerAddr = owner.String()
	if len(provider) > 0 {
		if err := sdk.ValidateAccAddress(provider); err != nil {
			return ctypes.ResultTx{}, err
		}
		providerAddr = provider
	}

	msg := &MsgDisableServiceBinding{
		ServiceName: serviceName,
		Provider:    providerAddr,
		Owner:       owner.String(),
	}
	return s.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

// EnableServiceBinding enables the specified service binding
func (s serviceClient) EnableServiceBinding(serviceName, provider string, deposit sdk.DecCoins, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	owner, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	var providerAddr = owner.String()
	if len(provider) > 0 {
		if err := sdk.ValidateAccAddress(provider); err != nil {
			return ctypes.ResultTx{}, err
		}
		providerAddr = provider
	}

	amt, err := s.ToMinCoin(deposit...)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	msg := &MsgEnableServiceBinding{
		ServiceName: serviceName,
		Provider:    providerAddr,
		Deposit:     amt,
		Owner:       owner.String(),
	}
	return s.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

//InvokeService is responsible for invoke a new service and callback `handler`
func (s serviceClient) InvokeService(request InvokeServiceRequest, baseTx sdk.BaseTx) (string, ctypes.ResultTx, error) {
	consumer, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return "", ctypes.ResultTx{}, err
	}

	var providers []string
	for _, provider := range request.Providers {
		if err := sdk.ValidateAccAddress(provider); err != nil {
			return "", ctypes.ResultTx{}, err
		}
		providers = append(providers, provider)
	}

	amt, err := s.ToMinCoin(request.ServiceFeeCap...)
	if err != nil {
		return "", ctypes.ResultTx{}, err
	}

	msg := &MsgCallService{
		ServiceName:       request.ServiceName,
		Providers:         providers,
		Consumer:          consumer.String(),
		Input:             request.Input,
		ServiceFeeCap:     amt,
		Timeout:           request.Timeout,
		Repeated:          request.Repeated,
		RepeatedFrequency: request.RepeatedFrequency,
		RepeatedTotal:     request.RepeatedTotal,
	}

	//mode must be set to commit
	baseTx.Mode = sdk.Commit

	result, err := s.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return "", ctypes.ResultTx{}, err
	}

	reqCtxID, e := sdk.StringifyEvents(result.TxResult.Events).GetValue(sdk.EventTypeCreateContext, attributeKeyRequestContextID)
	if e != nil {
		return reqCtxID, result, e
	}

	if request.Callback == nil {
		return reqCtxID, result, nil
	}

	_, err = s.SubscribeServiceResponse(reqCtxID, request.Callback)
	return reqCtxID, result, err
}

func (s serviceClient) InvokeServiceResponse(req InvokeServiceResponseRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	provider, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	reqId := req.RequestId
	_, err = s.QueryServiceRequest(reqId)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	msg := &MsgRespondService{
		RequestId: reqId,
		Provider:  provider.String(),
		Result:    req.Result,
		Output:    req.Output,
	}

	return s.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (s serviceClient) SubscribeServiceResponse(reqCtxID string,
	callback InvokeCallback) (subscription sdk.Subscription, err error) {
	if len(reqCtxID) == 0 {
		return subscription, errors.New(fmt.Sprintf("reqCtxID %s should not be empty", reqCtxID))
	}

	builder := sdk.NewEventQueryBuilder().AddCondition(
		sdk.NewCond(sdk.EventTypeResponseService, attributeKeyRequestContextID).EQ(sdk.EventValue(reqCtxID)),
	)

	return s.SubscribeTx(builder, func(tx sdk.EventDataTx) {
		s.Logger().Debug(
			"consumer received response transaction sent by provider",
			"tx_hash", tx.Hash,
			"height", tx.Height,
			"reqCtxID", reqCtxID,
		)
		for _, msg := range tx.Tx.GetMsgs() {
			msg, ok := msg.(*MsgRespondService)
			if ok {
				reqCtxID2, _, _, _, err := splitRequestID(msg.RequestId)
				if err != nil {
					s.Logger().Error(
						"invalid requestID",
						"requestID", msg.RequestId,
						"errMsg", err.Error(),
					)
					continue
				}
				if reqCtxID2.String() == strings.ToUpper(reqCtxID) {
					callback(reqCtxID, msg.RequestId, msg.Output)
				}
			}
		}
		reqCtx, err := s.QueryRequestContext(reqCtxID)
		if err != nil || reqCtx.State == RequestContextStateToStringMap[COMPLETED] {
			_ = s.Unsubscribe(subscription)
		}
	})
}

// SetWithdrawAddress sets a new withdrawal address for the specified service binding
func (s serviceClient) SetWithdrawAddress(withdrawAddress string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	owner, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	if err := sdk.ValidateAccAddress(withdrawAddress); err != nil {
		return ctypes.ResultTx{}, errors.New(fmt.Sprintf("%s invalid address", withdrawAddress))
	}
	msg := &MsgSetWithdrawAddress{
		Owner:           owner.String(),
		WithdrawAddress: withdrawAddress,
	}
	return s.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

// RefundServiceDeposit refunds the deposit from the specified service binding
func (s serviceClient) RefundServiceDeposit(serviceName, provider string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	owner, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	if err := sdk.ValidateAccAddress(provider); err != nil {
		return ctypes.ResultTx{}, err
	}

	msg := &MsgRefundServiceDeposit{
		ServiceName: serviceName,
		Provider:    provider,
		Owner:       owner.String(),
	}
	return s.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

// StartRequestContext starts the specified request context
func (s serviceClient) StartRequestContext(requestContextID string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	consumer, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}
	msg := &MsgStartRequestContext{
		RequestContextId: requestContextID,
		Consumer:         consumer.String(),
	}
	return s.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

// PauseRequestContext suspends the specified request context
func (s serviceClient) PauseRequestContext(requestContextID string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	consumer, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}
	msg := &MsgPauseRequestContext{
		RequestContextId: requestContextID,
		Consumer:         consumer.String(),
	}
	return s.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

// KillRequestContext terminates the specified request context
func (s serviceClient) KillRequestContext(requestContextID string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	consumer, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}
	msg := &MsgKillRequestContext{
		RequestContextId: requestContextID,
		Consumer:         consumer.String(),
	}
	return s.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

// UpdateRequestContext updates the specified request context
func (s serviceClient) UpdateRequestContext(request UpdateRequestContextRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	consumer, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	for _, provider := range request.Providers {
		if err := sdk.ValidateAccAddress(provider); err != nil {
			return ctypes.ResultTx{}, err
		}
	}

	amt, err := s.ToMinCoin(request.ServiceFeeCap...)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	msg := &MsgUpdateRequestContext{
		RequestContextId:  request.RequestContextID,
		Providers:         request.Providers,
		ServiceFeeCap:     amt,
		Timeout:           request.Timeout,
		RepeatedFrequency: request.RepeatedFrequency,
		RepeatedTotal:     request.RepeatedTotal,
		Consumer:          consumer.String(),
	}
	return s.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

// WithdrawEarnedFees withdraws the earned fees to the specified provider
func (s serviceClient) WithdrawEarnedFees(provider string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	owner, err := s.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, err
	}

	var providerAddr = owner.String()
	if len(provider) > 0 {
		if err := sdk.ValidateAccAddress(provider); err != nil {
			return ctypes.ResultTx{}, err
		}
		providerAddr = provider
	}

	msg := &MsgWithdrawEarnedFees{
		Owner:    owner.String(),
		Provider: providerAddr,
	}
	return s.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

//SubscribeSingleServiceRequest is responsible for registering a single service handler
func (s serviceClient) SubscribeServiceRequest(serviceName string,
	callback RespondCallback,
	baseTx sdk.BaseTx) (subscription sdk.Subscription, err error) {
	provider, e := s.QueryAddress(baseTx.From, baseTx.Password)
	if e != nil {
		return sdk.Subscription{}, e
	}

	builder := sdk.NewEventQueryBuilder().AddCondition(
		sdk.NewCond(eventTypeNewBatchRequestProvider, attributeKeyServiceName).EQ(sdk.EventValue(serviceName)),
	).AddCondition(
		sdk.NewCond(eventTypeNewBatchRequestProvider, attributeKeyProvider).EQ(sdk.EventValue(provider.String())),
	)

	return s.SubscribeNewBlock(builder, func(block sdk.EventDataNewBlock) {
		msgs := s.GenServiceResponseMsgs(block.ResultEndBlock.Events, serviceName, provider, callback)
		if msgs == nil || len(msgs) == 0 {
			s.Logger().Error("no message created",
				"serviceName", serviceName,
				"provider", provider,
			)
		}
		if _, err = s.SendBatch(msgs, baseTx); err != nil {
			s.Logger().Error("provider respond failed", "errMsg", err.Error())
		}
	})
}

// QueryDefinition return a service definition of the specified name
func (s serviceClient) QueryServiceDefinition(serviceName string) (QueryServiceDefinitionResponse, error) {
	conn, err := s.GenConn()

	if err != nil {
		return QueryServiceDefinitionResponse{}, err
	}

	resp, err := NewQueryClient(conn).Definition(
		context.Background(),
		&QueryDefinitionRequest{ServiceName: serviceName},
	)
	if err != nil {
		return QueryServiceDefinitionResponse{}, err
	}

	return resp.ServiceDefinition.Convert().(QueryServiceDefinitionResponse), nil
}

// QueryBinding return the specified service binding
func (s serviceClient) QueryServiceBinding(serviceName string, provider string) (QueryServiceBindingResponse, error) {
	conn, err := s.GenConn()

	if err != nil {
		return QueryServiceBindingResponse{}, err
	}

	if err := sdk.ValidateAccAddress(provider); err != nil {
		return QueryServiceBindingResponse{}, err
	}

	resp, err := NewQueryClient(conn).Binding(
		context.Background(),
		&QueryBindingRequest{
			ServiceName: serviceName,
			Provider:    provider,
		},
	)
	if err != nil {
		return QueryServiceBindingResponse{}, err
	}

	return resp.ServiceBinding.Convert().(QueryServiceBindingResponse), nil
}

// QueryBindings returns all bindings of the specified service
func (s serviceClient) QueryServiceBindings(serviceName string, pageReq *query.PageRequest) ([]QueryServiceBindingResponse, error) {
	conn, err := s.GenConn()

	if err != nil {
		return nil, err
	}

	resp, err := NewQueryClient(conn).Bindings(
		context.Background(),
		&QueryBindingsRequest{
			ServiceName: serviceName,
			Pagination:  pageReq,
		},
	)
	if err != nil {
		return nil, err
	}

	return serviceBindings(resp.ServiceBindings).Convert().([]QueryServiceBindingResponse), nil
}

// QueryRequest returns  the active request of the specified requestID
func (s serviceClient) QueryServiceRequest(requestID string) (QueryServiceRequestResponse, error) {
	conn, err := s.GenConn()

	if err != nil {
		return QueryServiceRequestResponse{}, err
	}

	resp, err := NewQueryClient(conn).Request(
		context.Background(),
		&QueryRequestRequest{RequestId: requestID},
	)

	if err == nil && !resp.Request.Empty() {
		return resp.Request.Convert().(QueryServiceRequestResponse), nil
	}

	//query service Request by block
	request, err := s.queryRequestByTxQuery(requestID)
	if err != nil {
		return QueryServiceRequestResponse{}, err
	}

	return request.Convert().(QueryServiceRequestResponse), nil
}

// QueryRequest returns all the active requests of the specified service binding
func (s serviceClient) QueryServiceRequests(serviceName string, provider string, pageReq *query.PageRequest) ([]QueryServiceRequestResponse, error) {
	conn, err := s.GenConn()

	if err != nil {
		return nil, err
	}

	if err := sdk.ValidateAccAddress(provider); err != nil {
		return nil, err
	}

	resp, err := NewQueryClient(conn).Requests(
		context.Background(),
		&QueryRequestsRequest{
			ServiceName: serviceName,
			Provider:    provider,
			Pagination:  pageReq,
		},
	)
	if err != nil {
		return nil, err
	}

	return requests(resp.Requests).Convert().([]QueryServiceRequestResponse), nil
}

// QueryRequestsByReqCtx returns all requests of the specified request context ID and batch counter
func (s serviceClient) QueryRequestsByReqCtx(reqCtxID string, batchCounter uint64, pageReq *query.PageRequest) ([]QueryServiceRequestResponse, error) {
	conn, err := s.GenConn()

	if err != nil {
		return nil, err
	}

	resp, err := NewQueryClient(conn).RequestsByReqCtx(
		context.Background(),
		&QueryRequestsByReqCtxRequest{
			RequestContextId: reqCtxID,
			BatchCounter:     batchCounter,
			Pagination:       pageReq,
		},
	)
	if err != nil {
		return nil, err
	}

	return requests(resp.Requests).Convert().([]QueryServiceRequestResponse), nil
}

// QueryResponse returns a response with the speicified request ID
func (s serviceClient) QueryServiceResponse(requestID string) (QueryServiceResponseResponse, error) {
	conn, err := s.GenConn()

	if err != nil {
		return QueryServiceResponseResponse{}, err
	}

	resp, err := NewQueryClient(conn).Response(
		context.Background(),
		&QueryResponseRequest{RequestId: requestID},
	)

	if err == nil {
		return resp.Response.Convert().(QueryServiceResponseResponse), nil
	}

	response, err := s.queryResponseByTxQuery(requestID)
	if err != nil {
		return QueryServiceResponseResponse{}, nil
	}

	return response.Convert().(QueryServiceResponseResponse), nil
}

// QueryResponses returns all responses of the specified request context and batch counter
func (s serviceClient) QueryServiceResponses(reqCtxID string, batchCounter uint64, pageReq *query.PageRequest) ([]QueryServiceResponseResponse, error) {
	conn, err := s.GenConn()

	if err != nil {
		return nil, err
	}

	resp, err := NewQueryClient(conn).Responses(
		context.Background(),
		&QueryResponsesRequest{
			RequestContextId: reqCtxID,
			BatchCounter:     batchCounter,
			Pagination:       pageReq,
		},
	)
	if err != nil {
		return nil, err
	}

	return responses(resp.Responses).Convert().([]QueryServiceResponseResponse), nil
}

// QueryRequestContext return the specified request context
func (s serviceClient) QueryRequestContext(reqCtxID string) (QueryRequestContextResp, error) {
	conn, err := s.GenConn()

	if err != nil {
		return QueryRequestContextResp{}, err
	}

	resp, err := NewQueryClient(conn).RequestContext(
		context.Background(),
		&QueryRequestContextRequest{RequestContextId: reqCtxID},
	)
	if err == nil && !resp.RequestContext.Empty() {
		return resp.RequestContext.Convert().(QueryRequestContextResp), nil
	}

	reqCtx, err := s.queryRequestContextByTxQuery(reqCtxID)
	if err != nil {
		return QueryRequestContextResp{}, err
	}

	return reqCtx.Convert().(QueryRequestContextResp), nil
}

//QueryFees return the earned fees for a provider
func (s serviceClient) QueryFees(provider string) (sdk.Coins, error) {
	if err := sdk.ValidateAccAddress(provider); err != nil {
		return nil, err
	}

	conn, err := s.GenConn()

	if err != nil {
		return nil, err
	}

	res, err := NewQueryClient(conn).EarnedFees(
		context.Background(),
		&QueryEarnedFeesRequest{Provider: provider},
	)
	if err != nil {
		return nil, err
	}
	return res.Fees, nil
}

func (s serviceClient) QueryParams() (QueryParamsResp, error) {
	conn, err := s.GenConn()

	if err != nil {
		return QueryParamsResp{}, err
	}

	res, err := NewQueryClient(conn).Params(
		context.Background(),
		&QueryParamsRequest{},
	)
	if err != nil {
		return QueryParamsResp{}, err
	}

	return res.Params.Convert().(QueryParamsResp), nil
}

func (s serviceClient) GenServiceResponseMsgs(events sdk.StringEvents, serviceName string,
	provider sdk.AccAddress,
	handler RespondCallback) (msgs []sdk.Msg) {

	var ids []string
	for _, e := range events {
		if e.Type != eventTypeNewBatchRequestProvider {
			continue
		}
		attributes := sdk.Attributes(e.Attributes)
		svcName := attributes.GetValue(attributeKeyServiceName)
		prov := attributes.GetValue(attributeKeyProvider)
		if svcName == serviceName && prov == provider.String() {
			reqIDsStr := attributes.GetValue(attributeKeyRequests)
			var idsTemp []string
			if err := json.Unmarshal([]byte(reqIDsStr), &idsTemp); err != nil {
				s.Logger().Error(
					"service request don't exist",
					attributeKeyRequestID, reqIDsStr,
					attributeKeyServiceName, serviceName,
					attributeKeyProvider, provider.String(),
					"errMsg", err.Error(),
				)
				return
			}
			ids = append(ids, idsTemp...)
		}
	}

	for _, reqID := range ids {
		request, err := s.QueryServiceRequest(reqID)
		if err != nil {
			s.Logger().Error(
				"service request don't exist",
				attributeKeyRequestID, reqID,
				attributeKeyServiceName, serviceName,
				attributeKeyProvider, provider.String(),
				"errMsg", err.Error(),
			)
			continue
		}
		//check again
		providerStr := provider.String()
		if providerStr == request.Provider && request.ServiceName == serviceName {
			output, result := handler(request.RequestContextID, reqID, request.Input)
			msgs = append(msgs, &MsgRespondService{
				RequestId: reqID,
				Provider:  providerStr,
				Output:    output,
				Result:    result,
			})
		}
	}
	return msgs
}
