package server

import (
	"github.com/dollarkillerx/PolygonNftDemo/internal/middleware"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func (s *Server) router() {
	s.app.Use(gin.Logger())

	v1 := s.app.Group("/api/v1")
	v1.POST("airdrop", middleware.MaxAllowed(3), s.airdrop)             // 空投
	v1.GET("address/:address/nfts", s.getUserNFT)                       // 获取某个地址下面所有的nft
	v1.GET("nft_holding_account/:erc721_token_id", s.nftHoldingAccount) // nft持有账户
	v1.GET("nft_holder", s.nftHolder)                                   // nft持有者

	s.app.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(200)
		content, err := ioutil.ReadFile("./polygon_demo/dist/index.html")
		if err != nil {
			c.Writer.WriteHeader(404)
			c.Writer.WriteString("Not Found")
			return
		}
		_, _ = c.Writer.Write(content)
		c.Writer.Header().Add("Accept", "text/html")
		c.Writer.Flush()
	})

	s.app.StaticFile("vite.svg", "./polygon_demo/dist/vite.svg")
	s.app.StaticFS("/assets", http.Dir("./polygon_demo/dist/assets"))
}
