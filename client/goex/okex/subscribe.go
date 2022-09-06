package okex

import (
	"context"
)

type OkexSubscribe struct {
	Channel string `json:"channel,omitempty"`
	InstId  string `json:"instId,omitempty"`
}

type OkexSubscribeReq struct {
	Op   string        `json:"op,omitempty"`
	Args []interface{} `json:"args,omitempty"`
}

type OkexSubscribeResp struct {
	Event string        `json:"event,omitempty"`
	Arg   OkexSubscribe `json:"args,omitempty"`
}

type SubscribeService struct {
	ExName string
	params []interface{}
	c      *Client
}

func (s *SubscribeService) Subscribe(x string) {
	s.params = append(s.params, x)
}

func (s *SubscribeService) SubscribeOk(channel, instId string) {
	x := OkexSubscribe{
		Channel: channel,
		InstId:  instId,
	}
	s.params = append(s.params, x)
}

func (s *SubscribeService) Do(ctx context.Context, ID uint) error {
	//msg := &OkexSubscribeReq{
	//	Op:   "subscribe",
	//	Args: s.params,
	//}
	//data, err := json.Marshal(msg)
	//if err != nil {
	//	return err
	//}
	//s.c.Send(data)
	return nil
}
