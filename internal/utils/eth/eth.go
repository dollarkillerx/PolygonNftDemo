package eth

import (
	"github.com/dollarkillerx/PolygonNftDemo/internal/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"log"
)

type eth struct {
	PrivateKey      string
	contractAddress string
	chainAddress    string
	EthClient       *ethclient.Client
	Erc721          *Erc721
}

var Eth *eth

func init() {
	Eth = &eth{
		PrivateKey:      conf.CONF.Token,
		contractAddress: "0xa10A5e7E2b2CC6fcaF4e74B00a163B10ee060eBf",
		chainAddress:    "https://rpc-mainnet.maticvigil.com",
	}

	dial, err := ethclient.Dial(Eth.chainAddress)
	if err != nil {
		log.Fatalln(err)
	}
	Eth.EthClient = dial

	constant := common.HexToAddress(Eth.contractAddress)
	erc721, err := NewErc721(constant, dial)
	if err != nil {
		log.Fatalln(err)
	}

	Eth.Erc721 = erc721
}
