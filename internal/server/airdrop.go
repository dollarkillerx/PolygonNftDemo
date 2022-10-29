package server

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/dollarkillerx/PolygonNftDemo/internal/pkg/errs"
	"github.com/dollarkillerx/PolygonNftDemo/internal/pkg/request"
	"github.com/dollarkillerx/PolygonNftDemo/internal/pkg/response"
	"github.com/dollarkillerx/PolygonNftDemo/internal/utils/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

func (s *Server) airdrop(ctx *gin.Context) {
	var (
		input request.Airdrop
		err   error
	)

	if err = ctx.ShouldBind(&input); err != nil {
		response.Return(ctx, errs.BadRequest)
		return
	}

	if !eth.CheckEthAddress(input.Address) {
		response.Return(ctx, errs.BadRequest)
		return
	}

	privateKey, err := crypto.HexToECDSA(eth.Eth.PrivateKey)
	if err != nil {
		response.Return(ctx, errs.SystemError)
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		response.Return(ctx, errs.SystemError)
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := eth.Eth.EthClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		response.Return(ctx, errs.SystemError)
		return
	}

	gasPrice, err := eth.Eth.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		response.Return(ctx, errs.SystemError)
		return
	}

	chainID, err := eth.Eth.EthClient.NetworkID(context.Background())
	if err != nil {
		log.Println(err)
		response.Return(ctx, errs.SystemError)
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Println(err)
		response.Return(ctx, errs.SystemError)
		return
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	mint, err := eth.Eth.Erc721.SafeMint(auth, common.HexToAddress(input.Address), "hSEh9Yc.jpeg")
	if err != nil {
		log.Println(err)
		response.Return(ctx, errs.SystemError)
		return
	}

	log.Println(mint.Hash().String())

	response.Return(ctx, nil)
}
