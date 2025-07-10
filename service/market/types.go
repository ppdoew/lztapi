package market

import "LztApi/lztapi/client"

type Market struct {
	cli *client.LztClient
}

func SetMarket(cli *client.LztClient) *Market {
	return &Market{cli: cli}
}
