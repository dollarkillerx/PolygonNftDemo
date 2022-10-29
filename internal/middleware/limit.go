package middleware

import (
	"github.com/dollarkillerx/PolygonNftDemo/internal/pkg/errs"
	"github.com/dollarkillerx/PolygonNftDemo/internal/pkg/response"
	"github.com/dollarkillerx/PolygonNftDemo/internal/utils"
	"github.com/gin-gonic/gin"

	"log"
)

// MaxAllowed 限制每秒请求数量
func MaxAllowed(limitValue int64) func(c *gin.Context) {
	limiter := utils.NewLimiter(limitValue)
	log.Println("limiter.SetMax:", limitValue)
	// 返回限流逻辑
	return func(c *gin.Context) {
		if !limiter.Ok() {
			response.Return(c, errs.Speeding) // //超过每秒200，就返回503错误码
			return
		}
		c.Next()
	}
}
