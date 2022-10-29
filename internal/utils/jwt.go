package utils

import (
	"github.com/dollarkillerx/PolygonNftDemo/internal/conf"
	"github.com/dollarkillerx/PolygonNftDemo/internal/pkg/enum"
	"github.com/dollarkillerx/PolygonNftDemo/internal/pkg/request"
	"github.com/dollarkillerx/jwt"
	"github.com/gin-gonic/gin"
)

var JWT *jwt.JWT

func InitJWT() {
	JWT = jwt.NewJwt(conf.CONF.JWTToken)
}

// GetAuthModel GetAuthModel
func GetAuthModel(ctx *gin.Context) request.AuthJWT {
	get, exists := ctx.Get(enum.AuthModel.String())
	if !exists {
		panic("what fuck JWTToken is not exists")
	}

	model, ok := get.(request.AuthJWT)
	if !ok {
		panic("what fuck JWTToken is not exists 2")
	}

	return model
}
