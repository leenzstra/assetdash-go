package client

import (
	"github.com/go-resty/resty/v2"
)

type AssetDashOption func(*AssetDashClient)

func WithBaseUrl(url string) AssetDashOption {
	return func(ad *AssetDashClient) {
		ad.Endpoints.BaseURL = url
	}
}

func WithHttpClient(client *resty.Client) AssetDashOption {
	return func(ad *AssetDashClient) {
		ad.Client = client
	}
}

func WithUserAgent(ua string) AssetDashOption {
	return func(ad *AssetDashClient) {
		ad.Client.SetHeader("User-Agent", ua)
	}
}
