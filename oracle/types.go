package oracle

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/errors"
)

const (
	ModuleName = "oracle"
)

var (
	_ sdk.Msg = &MsgCreateFeed{}
	_ sdk.Msg = &MsgStartFeed{}
	_ sdk.Msg = &MsgPauseFeed{}
	_ sdk.Msg = &MsgEditFeed{}
)

func (m MsgCreateFeed) ValidateBasic() error {
	if len(m.FeedName) == 0 {
		return errors.Wrap(ErrValidateBasic, "feedName missing")
	}

	if len(m.Providers) == 0 {
		return errors.Wrap(ErrValidateBasic, "providers missing")
	}

	if len(m.ServiceName) == 0 {
		return errors.Wrap(ErrValidateBasic, "serviceName missing")
	}

	if len(m.AggregateFunc) == 0 {
		return errors.Wrap(ErrValidateBasic, "aggregateFunc missing")
	}

	if len(m.ValueJsonPath) == 0 {
		return errors.Wrap(ErrValidateBasic, "valueJsonPath missing")
	}

	return nil
}

func (m MsgCreateFeed) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Creator)}
}

func (m MsgStartFeed) ValidateBasic() error {
	if len(m.FeedName) == 0 {
		return errors.Wrap(ErrValidateBasic, "feedName missing")
	}
	return nil
}

func (m MsgStartFeed) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Creator)}
}

func (m MsgPauseFeed) ValidateBasic() error {
	if len(m.FeedName) == 0 {
		return errors.Wrap(ErrValidateBasic, "feedName missing")
	}
	return nil
}

func (m MsgPauseFeed) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Creator)}
}

func (m MsgEditFeed) ValidateBasic() error {
	if len(m.FeedName) == 0 {
		return errors.Wrap(ErrValidateBasic, "feedName missing")
	}
	return nil
}

func (m MsgEditFeed) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Creator)}
}

func (f FeedContext) Convert() interface{} {
	return QueryFeedResp{
		Feed:              f.Feed,
		ServiceName:       f.ServiceName,
		Providers:         f.Providers,
		Input:             f.Input,
		Timeout:           f.Timeout,
		ServiceFeeCap:     f.ServiceFeeCap,
		RepeatedFrequency: f.RepeatedFrequency,
		ResponseThreshold: f.ResponseThreshold,
		State:             f.State,
	}
}

type Feeds []FeedContext

func (fs Feeds) Convert() interface{} {
	var frs []QueryFeedResp
	for _, f := range fs {
		frs = append(frs, QueryFeedResp{
			Feed:              f.Feed,
			ServiceName:       f.ServiceName,
			Providers:         f.Providers,
			Input:             f.Input,
			Timeout:           f.Timeout,
			ServiceFeeCap:     f.ServiceFeeCap,
			RepeatedFrequency: f.RepeatedFrequency,
			ResponseThreshold: f.ResponseThreshold,
			State:             f.State,
		})
	}
	return frs
}

type FeedValues []FeedValue

func (fvs FeedValues) Convert() interface{} {
	var fvrs []QueryFeedValueResp
	for _, fv := range fvs {
		fvrs = append(fvrs, QueryFeedValueResp{
			Data:      fv.Data,
			Timestamp: fv.Timestamp,
		})
	}
	return fvrs
}
