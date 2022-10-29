package server

import (
	"github.com/dollarkillerx/PolygonNftDemo/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Server) router() {
	s.app.Use(gin.Logger())

	v1 := s.app.Group("/api/v1")
	v1.POST("airdrop", middleware.MaxAllowed(3), s.airdrop) // 空投
	v1.GET("address/:address/nfts", s.getUserNFT)           // 获取某个地址下面所有的nft
	v1.GET("nft_holding_account/:id", s.nftHoldingAccount)  // nft持有账户
	v1.GET("nft_holder", s.nftHolder)                       // nft持有者
}
