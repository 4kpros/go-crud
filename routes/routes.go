package routes

import (
	"github.com/4kpros/go-crud/middleware"
	"github.com/gin-gonic/gin"
)

func GET(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.GET(endpoint, middleware.SecureAPIKeyHandler(handler, requiredAuth))
}

func POST(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.POST(endpoint, middleware.SecureAPIKeyHandler(handler, requiredAuth))
}

func PUT(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.PUT(endpoint, middleware.SecureAPIKeyHandler(handler, requiredAuth))
}

func PATCH(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.PATCH(endpoint, middleware.SecureAPIKeyHandler(handler, requiredAuth))
}

func DELETE(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.DELETE(endpoint, middleware.SecureAPIKeyHandler(handler, requiredAuth))
}
