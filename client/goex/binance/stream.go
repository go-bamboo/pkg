package binance

func NewBinanceStream(apiKey, secretKey string, streams []string) *Client {
	url := "wss://stream.binance.com:9443/stream?streams="
	for _, stream := range streams {
		url = url + stream + "/"
	}
	url = url[:len(url)-1]
	return New(APIKey(apiKey), SecretKey(secretKey), BaseURL(url))
}
