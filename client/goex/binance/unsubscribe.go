package binance

import (
	"context"
	"encoding/json"
)

type UnsubscribeReq struct {
	Method string        `json:"method,omitempty"`
	Params []interface{} `json:"params,omitempty"`
	ID     int64         `json:"id,omitempty"`
}

type UnsubscribeResp struct {
	Result []string `json:"result,omitempty"`
	ID     int64    `json:"id,omitempty"`
	Code   int      `json:"code,omitempty"`
	Msg    int      `json:"msg,omitempty"`
}

type UnsubscribeService struct {
	req SubscribeReq
	C   *Client
}

func (sub *UnsubscribeService) Subscribe(x string) {
	sub.req.Params = append(sub.req.Params, x)
}

func (sub *UnsubscribeService) Do(ctx context.Context, ID uint) error {
	sub.req.Method = "UNSUBSCRIBE"
	sub.C.RpcId = sub.C.RpcId + 1
	data, err := json.Marshal(sub.req)
	if err != nil {
		return err
	}
	sub.C.Send(data, sub.C.RpcId, sub)
	return nil
}

func (s *UnsubscribeService) Unmarshal(data []byte) (interface{}, error) {
	var ret UnsubscribeResp
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	delete(s.C.unpack, ret.ID)
	return &BinanceBody{Body: &ret}, nil
}
