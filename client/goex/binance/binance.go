package binance

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

const Binance = "Binance"

func (c *Client) UnpackTrade(data []byte) (interface{}, error) {
	k := gjson.Get(string(data), "id")
	if k.Type == gjson.Number {
		id := k.Int()
		unpack := c.unpack[id]
		return unpack.Unmarshal(data)
	}
	e := gjson.Get(string(data), "e")
	if e.Type == gjson.String {
		if e.String() == "aggTrade" {
			var ret AggTrade
			if err := json.Unmarshal(data, &ret); err != nil {
				return nil, err
			}
			return &BinanceBody{Body: &ret}, nil
		} else if e.String() == "trade" {
			var ret Trade
			if err := json.Unmarshal(data, &ret); err != nil {
				return nil, err
			}
			return &BinanceBody{Body: &ret}, nil
		} else if e.String() == "kline" {
			var ret Kline
			if err := json.Unmarshal(data, &ret); err != nil {
				return nil, err
			}
			return &BinanceBody{Body: &ret}, nil
		} else if e.String() == "24hrTicker" {
			var ret Ticker
			if err := json.Unmarshal(data, &ret); err != nil {
				return nil, err
			}
			return &BinanceBody{Body: &ret}, nil
		} else if e.String() == "depthUpdate" {
			var ret DepthUpdate
			if err := json.Unmarshal(data, &ret); err != nil {
				return nil, err
			}
			return &BinanceBody{Body: &ret}, nil
		}
	}
	return nil, ErrNotSupportType("not support")
}
