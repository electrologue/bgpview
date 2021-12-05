// Package bgpview contains an BGPView API client.
package bgpview

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

const defaultBaseURL = "https://api.bgpview.io"

// Client a BGPView API client.
type Client struct {
	baseURL *url.URL

	HTTPClient *http.Client
}

// NewClient creates a new Client.
func NewClient() *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	return &Client{
		baseURL:    baseURL,
		HTTPClient: &http.Client{Timeout: 5 * time.Second},
	}
}

// GetASN gets ASN.
func (c Client) GetASN(ctx context.Context, asNumber int) (*ASNInfo, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "asn", strconv.Itoa(asNumber)))
	if err != nil {
		return nil, err
	}

	var apiResp ASNInfo
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

// GetASNPrefixes gets ASN prefixes.
func (c Client) GetASNPrefixes(ctx context.Context, asNumber int) (*ASNPrefixesInfo, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "asn", strconv.Itoa(asNumber), "prefixes"))
	if err != nil {
		return nil, err
	}

	var apiResp ASNPrefixesInfo
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

// GetASNPeers gets ASN peers.
func (c Client) GetASNPeers(ctx context.Context, asNumber int) (*ASNPeersInfo, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "asn", strconv.Itoa(asNumber), "peers"))
	if err != nil {
		return nil, err
	}

	var apiResp ASNPeersInfo
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

// GetASNUpstreams gets ASN upstreams.
func (c Client) GetASNUpstreams(ctx context.Context, asNumber int) (*ASNUpstreamsInfo, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "asn", strconv.Itoa(asNumber), "upstreams"))
	if err != nil {
		return nil, err
	}

	var apiResp ASNUpstreamsInfo
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

// GetASNDownstreams gets ASN downstreams.
func (c Client) GetASNDownstreams(ctx context.Context, asNumber int) (*ASNDownstreamsInfo, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "asn", strconv.Itoa(asNumber), "downstreams"))
	if err != nil {
		return nil, err
	}

	var apiResp ASNDownstreamsInfo
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

// GetASNIxs gets ASN IXs.
func (c Client) GetASNIxs(ctx context.Context, asNumber int) (*ASNIxsInfo, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "asn", strconv.Itoa(asNumber), "ixs"))
	if err != nil {
		return nil, err
	}

	var apiResp ASNIxsInfo
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

// GetPrefix gets Prefix.
func (c Client) GetPrefix(ctx context.Context, ipAddress string, cidr int) (*PrefixInfo, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "prefix", ipAddress, strconv.Itoa(cidr)))
	if err != nil {
		return nil, err
	}

	var apiResp PrefixInfo
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

// GetIP gets IP.
func (c Client) GetIP(ctx context.Context, ipAddress string) (*IPInfo, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "ip", ipAddress))
	if err != nil {
		return nil, err
	}

	var apiResp IPInfo
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

// GetIX gets IX.
func (c Client) GetIX(ctx context.Context, ixID int) (*IXInfo, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "ix", strconv.Itoa(ixID)))
	if err != nil {
		return nil, err
	}

	var apiResp IXInfo
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

// GetSearch searches resources by ASN, IP, Prefix, Name, Description.
func (c Client) GetSearch(ctx context.Context, term string) (*SearchInfo, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "search"))
	if err != nil {
		return nil, err
	}

	query := endpoint.Query()
	query.Set("query_term", term)
	endpoint.RawQuery = query.Encode()

	var apiResp SearchInfo
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

func (c Client) do(ctx context.Context, endpoint *url.URL, data interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return err
	}

	req.Header.Set("accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%d: %s", resp.StatusCode, string(data))
	}

	return json.NewDecoder(resp.Body).Decode(data)
}
