package lztapi

import (
	"github.com/ppdoew/lztapi/client"
	"github.com/ppdoew/lztapi/service/market"
)

func CreateClient(token string) *client.LztClient {
	return client.NewSession(token)
}

func ModuleMarket(cli *client.LztClient) *market.Market {
	return market.SetMarket(cli)
}
