package handler

import (
	// "net/http"

	"github.com/Beksultan15/project_go/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
    return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	 {
		auth.POST("/signup",h.signup)
		auth.POST("/signin",h.signin)
		// auth.GET("/logout",)
	 }

	 search := router.Group("/products")
	 {
		search.GET("/search/",h.Search)

	 }

	 filter := router.Group("/products")
	 {
		filter.GET("/filter/lte=/:lte&gte=/:gte",h.filter)

	 }

	

	 return router
}