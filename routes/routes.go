package routes

import (
	"github.com/4kpros/go-crud/middleware"
	"github.com/gin-gonic/gin"
)

func GET(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.GET(endpoint, middleware.SecurityAPIKeyHandler(handler, requiredAuth))
}

func POST(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.POST(endpoint, middleware.SecurityAPIKeyHandler(handler, requiredAuth))
}

func PUT(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.PUT(endpoint, middleware.SecurityAPIKeyHandler(handler, requiredAuth))
}

func PATCH(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.PATCH(endpoint, middleware.SecurityAPIKeyHandler(handler, requiredAuth))
}

func DELETE(r *gin.Engine, endpoint string, handler gin.HandlerFunc, requiredAuth bool) {
	r.DELETE(endpoint, middleware.SecurityAPIKeyHandler(handler, requiredAuth))
}
