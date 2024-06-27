package router

import (
	"github.com/4kpros/go-api/common/middleware"
	"github.com/gin-gonic/gin"
)

func GET(router *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	router.GET(endpoint, middleware.SecureAPIKeyHandler(handler, requiredAuth))
}

func POST(router *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	router.POST(endpoint, middleware.SecureAPIKeyHandler(handler, requiredAuth))
}

func PUT(router *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	router.PUT(endpoint, middleware.SecureAPIKeyHandler(handler, requiredAuth))
}

func PATCH(router *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	router.PATCH(endpoint, middleware.SecureAPIKeyHandler(handler, requiredAuth))
}

func DELETE(router *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	router.DELETE(endpoint, middleware.SecureAPIKeyHandler(handler, requiredAuth))
}
