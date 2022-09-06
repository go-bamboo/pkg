package binance

type Option func(*options)

type options struct {
	apiKey    string
	secretKey string
	baseURL   string
}

func APIKey(key string) Option {
	return func(c *options) {
		c.apiKey = key
	}
}

func SecretKey(name string) Option {
	return func(c *options) {
		c.secretKey = name
	}
}

func BaseURL(path string) Option {
	return func(c *options) {
		c.baseURL = path
	}
}
