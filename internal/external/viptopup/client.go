package viptopup

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// New creates a new Client with required apiID, apiKey, and optional options.
func New(apiID, apiKey string, opts ...Option) *Client {
	sig := md5.Sum([]byte(apiID + apiKey))
	client := &Client{
		apiID:      apiID,
		apiKey:     apiKey,
		signature:  hex.EncodeToString(sig[:]),
		endpoint:   DefaultEndpoint,
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// Client is the VIPTopup API client.
type Client struct {
	apiID     string
	apiKey    string
	signature string
	endpoint  string

	httpClient Doer
	logger     Logger
	stats      Stats
}

// API endpoints
const (
	DefaultEndpoint = "https://vip-reseller.co.id/api"
)

// postForm is a helper to make a POST request with form data and decode the JSON response.
func (c *Client) postForm(ctx context.Context, endpoint string, params url.Values, out interface{}) error {
	form := params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBufferString(form))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status: " + resp.Status)
	}

	if err := json.Unmarshal(body, out); err != nil {
		return err
	}

	return nil
}

// buildParams returns a url.Values with the common key and sign parameters set.
func (c *Client) buildParams() url.Values {
	params := url.Values{}
	params.Set("key", c.apiKey)
	params.Set("sign", c.signature)
	return params
}

// Profile fetches the profile information.
func (c *Client) Profile(ctx context.Context) (*ProfileResponse, error) {
	params := c.buildParams()
	var resp ProfileResponse
	if err := c.postForm(ctx, c.endpoint+"/profile", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// OrderPrepaid places a prepaid order.
func (c *Client) OrderPrepaid(ctx context.Context, serviceCode, dataNo string) (*OrderPrepaidResponse, error) {
	params := c.buildParams()
	params.Set("type", "order")
	params.Set("service", serviceCode)
	params.Set("data_no", dataNo)
	var resp OrderPrepaidResponse
	if err := c.postForm(ctx, c.endpoint+"/prepaid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// StatusOrderPrepaid checks the status of a prepaid order.
func (c *Client) StatusOrderPrepaid(ctx context.Context, trxid string, limit *int) (*StatusOrderPrepaidResponse, error) {
	params := c.buildParams()
	params.Set("type", "status")
	params.Set("trxid", trxid)
	if limit != nil {
		params.Set("limit", fmt.Sprintf("%d", *limit))
	}
	var resp StatusOrderPrepaidResponse
	if err := c.postForm(ctx, c.endpoint+"/prepaid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ServicePrepaid fetches prepaid services.
func (c *Client) ServicePrepaid(ctx context.Context, filterType, filterValue *string) (*ServicePrepaidResponse, error) {
	params := c.buildParams()
	params.Set("type", "services")
	if filterType != nil {
		params.Set("filter_type", *filterType)
	}
	if filterValue != nil {
		params.Set("filter_value", *filterValue)
	}
	var resp ServicePrepaidResponse
	if err := c.postForm(ctx, c.endpoint+"/prepaid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// OrderGame places a game or streaming order.
func (c *Client) OrderGame(ctx context.Context, service, dataNo string, dataZone *string) (*OrderGameResponse, error) {
	params := c.buildParams()
	params.Set("type", "order")
	params.Set("service", service)
	params.Set("data_no", dataNo)
	if dataZone != nil {
		params.Set("data_zone", *dataZone)
	}
	var resp OrderGameResponse
	if err := c.postForm(ctx, c.endpoint+"/game-feature", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// StatusOrderGame checks the status of a game or streaming order.
func (c *Client) StatusOrderGame(ctx context.Context, trxid string, limit *int) (*StatusOrderGameResponse, error) {
	params := c.buildParams()
	params.Set("type", "status")
	params.Set("trxid", trxid)
	if limit != nil {
		params.Set("limit", fmt.Sprintf("%d", *limit))
	}
	var resp StatusOrderGameResponse
	if err := c.postForm(ctx, c.endpoint+"/game-feature", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ServiceGame fetches game and streaming services.
func (c *Client) ServiceGame(ctx context.Context, filterType, filterValue, filterStatus *string) (*ServiceGameResponse, error) {
	params := c.buildParams()
	params.Set("type", "services")
	if filterType != nil {
		params.Set("filter_type", *filterType)
	}
	if filterValue != nil {
		params.Set("filter_value", *filterValue)
	}
	if filterStatus != nil {
		params.Set("filter_status", *filterStatus)
	}
	var resp ServiceGameResponse
	if err := c.postForm(ctx, c.endpoint+"/game-feature", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Doer is an interface for http.Client to allow for dependency injection and testing.
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Stats is an interface for stats collection.
type Stats interface {
	Inc(name string, value int64)
}

// Logger is an interface for logging.
type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}
