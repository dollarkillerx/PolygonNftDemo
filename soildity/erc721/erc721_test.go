package erc721

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestErc721(t *testing.T) {
	dial, err := ethclient.Dial("https://rpc-mainnet.maticvigil.com")
	if err != nil {
		log.Fatalln(err)
	}

	// 合约地址
	constant := common.HexToAddress("0xa10A5e7E2b2CC6fcaF4e74B00a163B10ee060eBf")
	erc721, err := NewErc721(constant, dial)
	if err != nil {
		log.Fatalln(err)
	}

	of, err := erc721.OwnerOf(nil, big.NewInt(0))
	if err != nil {
		log.Fatalln(err)
	}

	printObj(of)
}

func printObj(i interface{}) {
	marshal, _ := json.Marshal(i)
	fmt.Println(string(marshal))
}
