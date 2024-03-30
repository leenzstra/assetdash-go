package client

import (
	"net/url"

	"github.com/EDDYCJY/fake-useragent"
	"github.com/go-resty/resty/v2"
)

type AssetDashClient struct {
	Endpoints *Endpoints
	Token     string
	Client    *resty.Client
}

func New(token string, endpoints *Endpoints, opts ...AssetDashOption) *AssetDashClient {
	ad := &AssetDashClient{
		Endpoints: endpoints,
		Token:     token,
		Client:    resty.New(),
	}

	ad.Client.SetHeaders(map[string]string{
		"Authorization": "Bearer " + ad.Token,
		"User-Agent":    browser.Random()},
	)

	ad.Client.Debug = true

	for _, option := range opts {
		option(ad)
	}

	return ad
}

func (ad *AssetDashClient) BuildURL(path ...string) (string, error) {
	return url.JoinPath(ad.Endpoints.BaseURL, path...)
}
