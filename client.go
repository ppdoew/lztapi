package lztapi

import (
	"LztApi/lztapi/client"
	"LztApi/lztapi/service/market"
)

func CreateClient(token string) *client.LztClient {
	return client.NewSession(token)
}

func ModuleMarket(cli *client.LztClient) *market.Market {
	return market.SetMarket(cli)
}
