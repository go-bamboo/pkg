package binance

import (
	"context"
	"encoding/json"
	"sync/atomic"
)

type SubscribeReq struct {
	Method string        `json:"method,omitempty"`
	Params []interface{} `json:"params,omitempty"`
	ID     int64         `json:"id,omitempty"`
}

type SubscribeResp struct {
	Result []string `json:"result,omitempty"`
	ID     int64    `json:"id,omitempty"`
	Code   int      `json:"code,omitempty"`
	Msg    int      `json:"msg,omitempty"`
}

type SubscribeService struct {
	req SubscribeReq
	C   *Client
}

func (s *SubscribeService) Subscribe(channel string) {
	s.req.Params = append(s.req.Params, channel)
}

func (s *SubscribeService) Do(ctx context.Context, ID uint) error {
	s.req.Method = "SUBSCRIBE"
	rpcId := atomic.AddInt64(&s.C.RpcId, 1)
	s.req.ID = rpcId
	data, err := json.Marshal(s.req)
	if err != nil {
		return err
	}
	s.C.Send(data, s.req.ID, s)
	return nil
}

func (s *SubscribeService) Unmarshal(data []byte) (interface{}, error) {
	var ret SubscribeResp
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	delete(s.C.unpack, ret.ID)
	return &BinanceBody{Body: &ret}, nil
}
