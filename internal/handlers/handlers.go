package handlers

import (
	"github.com/fishmanDK/miet_project/internal/service"
	"github.com/fishmanDK/miet_project/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handlers struct{
	service *service.Service
	log logger.Logger
}

func NewHandlers(service *service.Service) *Handlers{
	return &Handlers{
		service: service,
	}
}

func (h *Handlers) InitRouts() *gin.Engine {
	router := gin.New()

	auth := router.Group("/client")
	{
		auth.GET("/sign-in", h.signIn) //TODO
		auth.POST("/sign-up", h.signUp)
	}

	router.GET("/store", h.GetStore)//TODO
	router.POST("/store", h.CreateStore)
	router.DELETE("/store", h.DeleteStore) //TODO

	router.GET("/user", h.GetUser) //TODO
	router.PATCH("/user", h.UpdateUser) //TODO
	router.DELETE("/user", h.DeleteUser) //TODO

	router.POST("/cassette-availability", h.CreateCassetteAvailability)

	router.POST("/reservation", h.CreateReservation)

	cassette := router.Group("/cassette")
	{
		cassette.GET("", h.GetCassettes) //TODO
		cassette.GET("/:id", h.GetCassette) //TODO
		cassette.DELETE("", h.DeleteCassette) //TODO
		cassette.POST("", h.CreateCassette)
	}

	return router
}