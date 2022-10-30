package test

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/dollarkillerx/urllib"
)

func TestAlchemy(t *testing.T) {
	//address := "0xd1f84ec6652b315e77ebffa649f6742754574796"
	address2 := "0xa10A5e7E2b2CC6fcaF4e74B00a163B10ee060eBf"

	//kpl(address)
	kpl(address2)
}

func kpl(address string) {
	//input := map[string]interface{}{
	//	"jsonrpc": "2.0",
	//	"method":  "alchemy_getTokenBalances",
	//	"headers": map[string]interface{}{
	//		"Content-Type": "application/json",
	//	},
	//	"params": []string{
	//		address,
	//		"erc721",
	//		"erc1155",
	//	},
	//	"id": 42,
	//}

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
					"0xa10A5e7E2b2CC6fcaF4e74B00a163B10ee060eBf",
				},
				"withMetadata":     false,
				"excludeZeroValue": true,
				"maxCount":         "0x3e8",
				//"fromAddress":      address,
			},
		},
	}

	i, bytes, err := urllib.
		Post("https://polygon-mainnet.g.alchemy.com/v2/g9_dtmuJeERcETxHWLr2fBqgdz9qcie4").
		SetJsonObject(input).Byte()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(i)
	fmt.Println(string(bytes))

	os.WriteFile("alchemy.json", bytes, 00666)
}

func TestP2(t *testing.T) {
	r := "0x0000000000000000000000000000000000000000000000000000000000000001"
	val := r[2:]
	i, err := strconv.ParseInt(val, 16, 64)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(i)
}
