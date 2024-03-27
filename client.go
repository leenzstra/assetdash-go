package main

type AssetDash struct {
	baseUrl string
}

type AssetDashOption func(*AssetDash)

func New(options ...AssetDashOption) *AssetDash {
	ad := &AssetDash{}
	for _, option := range options {
		option(ad)
	}
	return ad
}

func (ad *AssetDash) WithBaseUrl(url string) AssetDashOption {
	return func(ad *AssetDash) {
		ad.baseUrl = url
	}

}
