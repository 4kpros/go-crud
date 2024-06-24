package router

import (
	"github.com/4kpros/go-crud/common/middlewares"
	"github.com/gin-gonic/gin"
)

func GET(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.GET(endpoint, middlewares.SecureAPIKeyHandler(handler, requiredAuth))
}

func POST(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.POST(endpoint, middlewares.SecureAPIKeyHandler(handler, requiredAuth))
}

func PUT(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.PUT(endpoint, middlewares.SecureAPIKeyHandler(handler, requiredAuth))
}

func PATCH(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.PATCH(endpoint, middlewares.SecureAPIKeyHandler(handler, requiredAuth))
}

func DELETE(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.DELETE(endpoint, middlewares.SecureAPIKeyHandler(handler, requiredAuth))
}
