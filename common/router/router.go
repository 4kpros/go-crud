package router

import (
	"github.com/4kpros/go-api/common/middleware"
	"github.com/gin-gonic/gin"
)

func GET(ginRouter *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	ginRouter.GET(endpoint, middleware.SecureAPIHandler(handler, requiredAuth))
}

func POST(ginRouter *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	ginRouter.POST(endpoint, middleware.SecureAPIHandler(handler, requiredAuth))
}

func PUT(ginRouter *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	ginRouter.PUT(endpoint, middleware.SecureAPIHandler(handler, requiredAuth))
}

func PATCH(ginRouter *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	ginRouter.PATCH(endpoint, middleware.SecureAPIHandler(handler, requiredAuth))
}

func DELETE(ginRouter *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	ginRouter.DELETE(endpoint, middleware.SecureAPIHandler(handler, requiredAuth))
}
