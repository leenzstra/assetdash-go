package main

import "net/http"

type AssetDashOption func(*AssetDash)

func WithBaseUrl(url string) AssetDashOption {
	return func(ad *AssetDash) {
		ad.baseUrl = url
	}
}

func WithHttpClient(client *http.Client) AssetDashOption {
	return func(ad *AssetDash) {
		ad.client = client
	}
}