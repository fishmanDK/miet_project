package handlers

import (
	"fmt"
	"net/http"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateReservation (c *gin.Context) {
    var input core.Reservation
	if err := c.BindJSON(&input); err != nil{
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return 
	}

	id, err := h.service.Reservation.CreateReservation(input)
	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return 
	}

	input.Id = id
	input.Status = "reserve"
	c.JSON(http.StatusOK, input)
}