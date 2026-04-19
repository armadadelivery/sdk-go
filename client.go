// Package armada is the official Go SDK for the Armada Automated Ordering
// API v2. It signs every request with HMAC-SHA256 and surfaces rate-limit
// headers on every response.
//
// Usage:
//
//	c := armada.NewClient(armada.Options{APIKey: key, APISecret: secret})
//	resp, err := c.Get(ctx, "/v2/wallet", nil)
//	defer resp.Body.Close()
package armada

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const defaultBase = "https://api.armadadelivery.com"

type Options struct {
	APIKey    string
	APISecret string
	// Defaults to https://api.armadadelivery.com (production). Use a
	// Test-mode API key to simulate deliveries without dispatching a real
	// driver. Pass https://sandbox.api.armadadelivery.com explicitly only
	// if you need the fully-isolated sandbox environment.
	BaseURL string
	Timeout time.Duration // default 30s
}

type Client struct {
	http      *http.Client
	baseURL   string
	apiKey    string
	apiSecret string
}

type RateLimit struct {
	Limit     *int
	Remaining *int
	ResetUnix *int64
}

type Response struct {
	*http.Response
	RateLimit RateLimit
}

func NewClient(o Options) *Client {
	if o.APIKey == "" || o.APISecret == "" {
		panic("armada.NewClient: APIKey + APISecret are required")
	}
	base := o.BaseURL
	if base == "" {
		base = defaultBase
	}
	timeout := o.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}
	return &Client{
		http:      &http.Client{Timeout: timeout},
		baseURL:   base,
		apiKey:    o.APIKey,
		apiSecret: o.APISecret,
	}
}

func (c *Client) sign(timestamp, method, path, body string) string {
	payload := fmt.Sprintf("%s.%s.%s.%s", timestamp, method, path, body)
	mac := hmac.New(sha256.New, []byte(c.apiSecret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}

// Do sends an HTTP request to the v2 API.
func (c *Client) Do(ctx context.Context, method, path string, query url.Values, body interface{}) (*Response, error) {
	var bodyBytes []byte
	var err error
	if body != nil {
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	full := c.baseURL + path
	if len(query) > 0 {
		full += "?" + query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, full, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	signedPath := path
	if len(query) > 0 {
		signedPath += "?" + query.Encode()
	}
	req.Header.Set("Authorization", "Key "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "armada-sdk-go")
	req.Header.Set("x-armada-timestamp", timestamp)
	req.Header.Set("x-armada-signature", c.sign(timestamp, method, signedPath, string(bodyBytes)))

	raw, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	return &Response{Response: raw, RateLimit: parseRateLimit(raw.Header)}, nil
}

// Get is sugar for Do(GET, path, query, nil).
func (c *Client) Get(ctx context.Context, path string, query url.Values) (*Response, error) {
	return c.Do(ctx, http.MethodGet, path, query, nil)
}

// PostJSON is sugar for Do(POST, path, nil, body).
func (c *Client) PostJSON(ctx context.Context, path string, body interface{}) (*Response, error) {
	return c.Do(ctx, http.MethodPost, path, nil, body)
}

func (c *Client) PutJSON(ctx context.Context, path string, body interface{}) (*Response, error) {
	return c.Do(ctx, http.MethodPut, path, nil, body)
}

func (c *Client) Delete(ctx context.Context, path string) (*Response, error) {
	return c.Do(ctx, http.MethodDelete, path, nil, nil)
}

// DecodeJSON reads + closes the response body and unmarshals into dst. It
// returns an error for 4xx/5xx responses carrying the decoded error body.
func DecodeJSON(resp *Response, dst interface{}) error {
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("armada API %d: %s", resp.StatusCode, string(raw))
	}
	if dst == nil {
		return nil
	}
	return json.Unmarshal(raw, dst)
}

func parseRateLimit(h http.Header) RateLimit {
	intp := func(s string) *int {
		if s == "" {
			return nil
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		return &n
	}
	int64p := func(s string) *int64 {
		if s == "" {
			return nil
		}
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil
		}
		return &n
	}
	return RateLimit{
		Limit:     intp(h.Get("X-RateLimit-Limit")),
		Remaining: intp(h.Get("X-RateLimit-Remaining")),
		ResetUnix: int64p(h.Get("X-RateLimit-Reset")),
	}
}
