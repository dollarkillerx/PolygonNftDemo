package server

import (
	"github.com/dollarkillerx/urllib"

	"log"
	"time"
)

func (s *Server) alchemy() {
	for {
		select {
		case <-time.After(time.Second):
			s.alchemyInternal()
		}
	}
}

type AlchemyResp struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		Transfers []struct {
			BlockNum        string      `json:"blockNum"`
			UniqueId        string      `json:"uniqueId"`
			Hash            string      `json:"hash"`
			From            string      `json:"from"`
			To              string      `json:"to"`
			Value           interface{} `json:"value"`
			Erc721TokenId   string      `json:"erc721TokenId"`
			Erc1155Metadata interface{} `json:"erc1155Metadata"`
			TokenId         string      `json:"tokenId"`
			Asset           interface{} `json:"asset"`
			Category        string      `json:"category"`
			RawContract     struct {
				Value   interface{} `json:"value"`
				Address string      `json:"address"`
				Decimal interface{} `json:"decimal"`
			} `json:"rawContract"`
		} `json:"transfers"`
	} `json:"result"`
}

func (s *Server) alchemyInternal() {
	contractAddresses := "0xa10A5e7E2b2CC6fcaF4e74B00a163B10ee060eBf"

	input := map[string]interface{}{
		"id":      1,
		"jsonrpc": "2.0",
		"method":  "alchemy_getAssetTransfers",
		"params": []map[string]interface{}{
			{
				"fromBlock": "0x0",
				"toBlock":   "latest",
				"category": []string{
					"erc721",
				},
				"contractAddresses": []string{
					contractAddresses,
				},
				"withMetadata":     false,
				"excludeZeroValue": true,
				"maxCount":         "0x3e8",
				//"fromAddress":      address,
			},
		},
	}

	var alchemyResp AlchemyResp
	err := urllib.
		Post("https://polygon-mainnet.g.alchemy.com/v2/g9_dtmuJeERcETxHWLr2fBqgdz9qcie4").
		SetJsonObject(input).SetTimeout(time.Second*10).FromJsonByCode(&alchemyResp, 200)
	if err != nil {
		log.Println(err)
		return
	}

	rdata := map[string]string{}
	for _, v := range alchemyResp.Result.Transfers {
		rdata[v.Erc721TokenId] = v.To
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache = rdata
}

func (s *Server) getCache() map[string]string {
	s.mu.Lock()
	defer s.mu.Unlock()

	var r = map[string]string{}
	for k, v := range s.cache {
		r[k] = v
	}

	return r
}
