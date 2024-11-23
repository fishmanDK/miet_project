package handlers

import (
	"fmt"
	"net/http"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateReservation(c *gin.Context) {
	var input core.Reservation
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := h.service.Reservation.CreateReservation(input)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handlers) DeleteReservation(c *gin.Context) {
	userID := c.GetInt("user_id")
	if userID == 0{
		fmt.Println("user_id = 0")
		return 
	}
	var input core.DeleteReservation
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := h.service.Reservation.DeleteReservation(userID, input.CassetteId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}


	c.JSON(http.StatusOK, nil)
}

func (h *Handlers) GetReservations(c *gin.Context) {
	userID := c.GetInt("user_id")
	if userID == 0{
		fmt.Println("user_id = 0")
		return 
	}

	reservations, err := h.service.Reservation.GetUserReservations(userID)
	if err != nil{
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, nil)
		return
	}

	for i := range reservations{
		reservations[i].ReservationDate = reservations[i].ReservationDate[:10]
	}

	err = h.tmpls.ExecuteTemplate(c.Writer, "reservations.html", struct {
		Reservations []core.Reservation
	}{
		Reservations: reservations,
	})
}