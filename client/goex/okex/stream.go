package okex

func NewOkexStream(apiKey, secretKey string, streamName string) *Client {
	// url := "wss://ws.okex.com:8443/ws/v5/public"
	// if true {
	// 	url = "wss://testnet.binance.vision/ws"
	// }
	// url = url + "/" + streamName
	url := "wss://wspap.okex.com:8443/ws/v5/public?brokerId=9999"
	return New(BaseURL(url))
}
