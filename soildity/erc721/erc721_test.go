package erc721

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
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

func TestErc721V2(t *testing.T) {
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

	// 加载您的私钥
	token := ""

	privateKey, err := crypto.HexToECDSA(token)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 读取我们应该用于帐户交易的随机数
	nonce, err := dial.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 獲取gas費用
	gasPrice, err := dial.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 獲取鏈id
	chainID, err := dial.NetworkID(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalln(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	mint, err := erc721.SafeMint(auth, common.HexToAddress("0xd1f84EC6652b315e77EbFFA649f6742754574796"), "hSEh9Yc.jpeg")
	if err != nil {
		log.Fatalln(err)
	}

	mint.Hash().String()
	printObj(mint)
}

func printObj(i interface{}) {
	marshal, _ := json.Marshal(i)
	fmt.Println(string(marshal))
}
