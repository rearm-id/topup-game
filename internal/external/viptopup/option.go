package viptopup

// Option pattern
type Option func(*Client)

func WithHTTPClient(httpClient Doer) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func WithEndpoint(endpoint string) Option {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

func WithLogger(logger Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

func WithStats(stats Stats) Option {
	return func(c *Client) {
		c.stats = stats
	}
}
