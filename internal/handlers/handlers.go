package handlers

import (
	"html/template"
	"log"
	"path/filepath"

	"github.com/fishmanDK/miet_project/internal/service"
	"github.com/fishmanDK/miet_project/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handlers struct{
	service *service.Service
	log logger.Logger
	tmpls *template.Template
}

func NewHandlers(service *service.Service, tmpls *template.Template, log logger.Logger) *Handlers{
	return &Handlers{
		service: service,
		tmpls: tmpls,
		log: log,
	}
}

func (h *Handlers) InitRouts() *gin.Engine {
	router := gin.New()

	absPath, err := filepath.Abs("./static")
    if err != nil {
        log.Fatalf("Ошибка при определении пути к статике: %v", err)
    }
    log.Println("Путь к статическим файлам:", absPath)

    router.Static("/static", absPath)

	router.GET("/", h.Authentication, h.Index)

	check := router.Group("/check", h.Authentication)
	{
		check.GET("", h.Check)
	}
	// check.Use(h.Authentication)

	auth := router.Group("/auth")
	{
		auth.GET("", h.RedirectSignIn)
		auth.GET("/sign-in", h.signInPage)
		auth.POST("/sign-in", h.signIn) //TODO
		auth.GET("/sign-up", h.signUpPage)
		auth.POST("/sign-up", h.signUp)
	}

	store := router.Group("/store", h.Authentication)
	{
		store.GET("", h.GetStores)//TODO
		store.GET("/:id", h.GetStore)
		store.POST("", h.CreateStore)
		store.DELETE("/:id", h.DeleteStore) //TODO
	}

	orders := router.Group("/orders", h.Authentication)
	{
		orders.GET("", h.GetOrders)
		orders.POST("", h.CreateOrder)
		orders.DELETE("/:id", h.DeleteOrder)
	}


	user := router.Group("/user", h.Authentication)
	{
		user.GET("", h.GetUser) //TODO
		user.PATCH("", h.UpdateUser) //TODO
		user.DELETE("", h.DeleteUser) //TODO
	}
	

	router.POST("/cassette-availability", h.CreateCassetteAvailability)

	reservation := router.Group("/reservations", h.Authentication)
	{
		reservation.GET("", h.GetReservations)
		reservation.POST("", h.CreateReservation)
		reservation.DELETE("", h.DeleteReservation)
	}
	

	cassette := router.Group("/cassette", h.Authentication)
	{
		
		cassette.GET("", h.GetCassettes) //TODO
		cassette.GET("/:id", h.GetCassette) //TODO
		cassette.GET("/details/:id", h.GetCassetteDetails) //TODO
		cassette.DELETE("/:id", h.DeleteCassette) //TODO
		cassette.POST("", h.CreateCassette)
	}

	// user.Use(h.Authentication)
	// reservation.Use(h.Authentication)
	// cassette.Use(h.Authentication)
	// store.Use(h.Authentication)

	return router
}