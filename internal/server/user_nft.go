package server

import (
	"github.com/dollarkillerx/PolygonNftDemo/internal/pkg/errs"
	"github.com/dollarkillerx/PolygonNftDemo/internal/pkg/response"
	"github.com/dollarkillerx/PolygonNftDemo/internal/utils/eth"
	"github.com/gin-gonic/gin"

	"log"
	"math/big"
)

func (s *Server) getUserNFT(ctx *gin.Context) {
	address := ctx.Param("address")

	if !eth.CheckEthAddress(address) {
		response.Return(ctx, errs.BadRequest)
		return
	}

	cache := s.getCache()
	var rdata []string
	for k, v := range cache {
		if v == address {
			rdata = append(rdata, k)
		}
	}

	response.Return(ctx, rdata)
}

func (s *Server) nftHoldingAccount(ctx *gin.Context) {
	erc721TokenId := ctx.Param("erc721_token_id")

	atoi, err := eth.Parse16TO10(erc721TokenId)
	if err != nil {
		response.Return(ctx, errs.BadRequest)
		return
	}

	of, err := eth.Eth.Erc721.OwnerOf(nil, big.NewInt(int64(atoi)))
	if err != nil {
		log.Println(err)
		response.Return(ctx, errs.SystemError)
		return
	}

	response.Return(ctx, of.String())
}

func (s *Server) nftHolder(ctx *gin.Context) {
	cache := s.getCache()
	response.Return(ctx, cache)
}
